package main

import (
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

	// massage channel
	ch := make(chan string)

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
					fmt.Println(string(buff))
					ch <- string(buff)
				}
			}
		}()

		go func() {
			for {
				message := <-ch
				// send message to client
				if len(connections) > 0 {
					for _, c := range connections {
						_, err = c.Write([]byte(message))
						if err != nil {
							continue
						}
					}
				}
			}
		}()
	}
}
