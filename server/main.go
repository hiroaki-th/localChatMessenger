package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const fileName string = "tmp_file"

func main() {

	// remove file if exist
	os.Remove(fileName)

	// listen
	// net.Listen(network, address string)
	listener, err := net.Listen("unix", fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("start server")
	fmt.Printf("listening at %s\n", listener.Addr())
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Printf("\n")

	conn, err := listener.Accept()
	for {
		if conn == nil {
			conn, _ = listener.Accept()
		}

		// wait for connection from client
		if err != nil {
			fmt.Printf("cannot connect to a client: %s\n", err)
			os.Exit(1)
		}

		// read from client
		buff := make([]byte, 64)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("cannot read: %s\n", err)
			conn.Close()
		}

		fmt.Printf("client:   %s\n", string(buff))

		// write message to client
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("you:   ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannot read what you input\n")
		}
		fmt.Printf("\n")

		// send message to client
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Printf("cannot send message to client: %s\n", err)
		}
	}
}
