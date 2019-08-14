package main

import (
	"./config"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	servAddr := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	conn, err := net.Dial("udp", servAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("<<< Client Send [", n, "] bytes")

		recvData := make([]byte, config.RECV_LEN)
		n, err = conn.Read(recvData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(">>> Client Recveived [", n, "] bytes")
	}
}
