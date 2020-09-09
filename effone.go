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

	fmt.Println(pkt.PacketID, pkt)
	if pkt.PacketID == f1.PacketIDCarStatus {
		var cspkt f1.CarStatusData
		cserr := binary.Read(buf, binary.LittleEndian, &cspkt)
		if cserr != nil {
			fmt.Println("CarStatus: binary read failed")
			return
		}
		fmt.Println(cspkt)
	} else if pkt.PacketID == f1.PacketIDCarTelemetry {
		var ctpkt f1.CarTelemetryData
		cterr := binary.Read(buf, binary.LittleEndian, &ctpkt)
		if cterr != nil {
			fmt.Println("CarTelemetry: binary read failed")
			return
		}
		fmt.Println(ctpkt)
	}
}
