package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	f1 "github.com/claesp/effone/libeffone"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func main() {
	go func() {
		addr, err := net.ResolveUDPAddr("udp", ":20777")
		if err != nil {
			fmt.Println("Wrong Address")
			return
		}

		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		fmt.Println("listening on", addr, "for F1 2020 telemetry")
		for {
			handleTelemetry(conn)
		}
	}()

	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	fmt.Println("listening on :8080 for web requests")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleTelemetry(conn *net.UDPConn) {
	readTelemetry(conn)
}

func readTelemetry(conn *net.UDPConn) {
	tmp := make([]byte, 4096)
	conn.Read(tmp)
	buf := bytes.NewBuffer(tmp)

	var pkt f1.PacketHeader
	berr := binary.Read(buf, binary.LittleEndian, &pkt)
	if berr != nil {
		fmt.Println("binary failure")
		return
	}

	parse(pkt.PacketID, tmp)
}

func parse(pid f1.PacketHeaderPacketID, data []byte) {
	buf := bytes.NewBuffer(data)
	switch pid {
	case f1.PacketIDMotion:
		var mpkt f1.PacketMotionData
		merr := binary.Read(buf, binary.LittleEndian, &mpkt)
		if merr != nil {
			fmt.Println("MotionData: binary read failed")
			return
		}
		fmt.Println(mpkt)
	case f1.PacketIDSession:
		var spkt f1.PacketSessionData
		serr := binary.Read(buf, binary.LittleEndian, &spkt)
		if serr != nil {
			fmt.Println("SessionData: binary read failed")
			return
		}
		fmt.Println(spkt)
	case f1.PacketIDLapData:
		var ldpkt f1.PacketLapData
		lderr := binary.Read(buf, binary.LittleEndian, &ldpkt)
		if lderr != nil {
			fmt.Println("LapData: binary read failed")
			return
		}
		fmt.Println(ldpkt)
	case f1.PacketIDEvent:
		var epkt f1.PacketEventData
		eerr := binary.Read(buf, binary.LittleEndian, &epkt)
		if eerr != nil {
			fmt.Println("EventData: binary read failed")
			return
		}
		fmt.Println(epkt)
	case f1.PacketIDParticipants:
		var ppkt f1.PacketParticipantsData
		perr := binary.Read(buf, binary.LittleEndian, &ppkt)
		if perr != nil {
			fmt.Println("ParticipantsData: binary read failed")
			return
		}
		fmt.Println(ppkt)
	case f1.PacketIDCarSetups:
		var cspkt f1.PacketCarSetupData
		cserr := binary.Read(buf, binary.LittleEndian, &cspkt)
		if cserr != nil {
			fmt.Println("CarSetupData: binary read failed")
			return
		}
		fmt.Println(cspkt)
	case f1.PacketIDCarTelemetry:
		var ctpkt f1.CarTelemetryData
		cterr := binary.Read(buf, binary.LittleEndian, &ctpkt)
		if cterr != nil {
			fmt.Println("CarTelemetryData: binary read failed")
			return
		}
		fmt.Println(ctpkt)
	case f1.PacketIDCarStatus:
		var cspkt f1.CarStatusData
		cserr := binary.Read(buf, binary.LittleEndian, &cspkt)
		if cserr != nil {
			fmt.Println("CarStatusData: binary read failed")
			return
		}
		fmt.Println(cspkt)
	case f1.PacketIDFinalClassification:
		var fcpkt f1.PacketFinalClassificationData
		fcerr := binary.Read(buf, binary.LittleEndian, &fcpkt)
		if fcerr != nil {
			fmt.Println("FinalClassificationData: binary read failed")
			return
		}
		fmt.Println(fcpkt)
	case f1.PacketIDLobbyInfo:
		var lipkt f1.PacketLobbyInfoData
		lierr := binary.Read(buf, binary.LittleEndian, &lipkt)
		if lierr != nil {
			fmt.Println("LobbyInfoData: binary read failed")
			return
		}
		fmt.Println(lipkt)
	default:
		fmt.Println("Unknown packet: %s", pid)
	}
}
