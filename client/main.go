package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// create connection to server
	conn, err := net.Dial("unix", "../server/tmp_file")
	if err != nil {
		fmt.Printf("cannot connect to server: %s", err)
		os.Exit(1)
	}
	fmt.Println("connect server")
	fmt.Printf("connecting to %s\n", conn.RemoteAddr().String())
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Printf("\n")

	// write message to serer
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("you:   ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannot read what you input\n")
		}
		fmt.Printf("\n")

		conn.Write([]byte(text))

		// receive response from server
		buf := make([]byte, 64)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("cannot read response: %s", err)
			os.Exit(1)
		}

		fmt.Printf("server:   %s\n", string(buf))
	}
}
