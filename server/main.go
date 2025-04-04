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

	// manage connection
	connections := make([]net.Conn, 0)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		connections = append(connections, conn)

		go func() {
			for {
				// read from client
				buff := make([]byte, 64)
				conn.Read(buff)
				if buff[0] != 0 {
					fmt.Printf("client:   %s\n", string(buff))
				}
			}
		}()

		// write message to client
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannot read what you input\n")
		}
		fmt.Printf("\n")

		// send message to client
		if len(connections) > 0 {
			for _, c := range connections {
				_, err = c.Write([]byte(text))
				if err != nil {
					fmt.Printf("cannot send message to client: %s\n", err)
				}
			}
		}
	}
}
