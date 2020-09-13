package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	f1 "github.com/claesp/effone/libeffone"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":20777")
	if err != nil {
		fmt.Println("Wrong Address")
		return
	}

	go func() {
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		fmt.Println("listening on", addr)
		for {
			handle(conn)
		}
	}()

	addr2, err := net.ResolveUDPAddr("udp", ":20778")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn2, err := net.ListenUDP("udp", addr2)
	if err != nil {
		fmt.Println(err)
	}
	defer conn2.Close()

	fmt.Println("listening on", addr2)
	for {
		handle(conn2)
	}
}

func handle(conn *net.UDPConn) {
	read(conn)
}

func read(conn *net.UDPConn) {
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
