package main

import (
	"./config"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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
		data := make([]byte, config.SERVER_RECV_LEN)
		n, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data[:n])
		fmt.Println("Server Recv: [", n, "]", strData)

		upper := strings.ToUpper(strData)
		n, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Server Send: [", n, "]", upper)
	}
}
