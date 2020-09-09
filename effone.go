package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	lf1 "github.com/claesp/effone/libeffone"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":20777")
	if err != nil {
		fmt.Println("Wrong Address")
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("listening on", addr)

	for {
		read(conn)
	}
}

func read(conn *net.UDPConn) {
	tmp := make([]byte, 1024)
	conn.Read(tmp)
	buf := bytes.NewBuffer(tmp)

	var pkt lf1.PacketHeader
	berr := binary.Read(buf, binary.LittleEndian, &pkt)
	if berr != nil {
		fmt.Println("binary failure")
		return
	}
	fmt.Println(pkt)
}
