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
		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > config.SERVER_RECV_LEN {
				toWrite = line[written : written+config.SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Client Send: [", n, "]", toWrite)

			msg := make([]byte, config.SERVER_RECV_LEN)
			n, err = conn.Read(msg)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Client Recv: [", n, "]", string(msg[:n]))
		}
	}
}
