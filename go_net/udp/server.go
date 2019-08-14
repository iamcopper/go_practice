package main

import (
	"./config"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server is listening on ", address)

	defer conn.Close()

	for {
		fmt.Println("--------------------------------")

		recvData := make([]byte, config.RECV_LEN)
		n, rAddr, err := conn.ReadFromUDP(recvData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(">>> Server received [", n, "] bytes")

		sendData := recvData[:n]
		n, err = conn.WriteToUDP(sendData, rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("<<< Server send [", n, "] bytes")
	}
}
