package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//get username
	username := "client"
	if len(os.Args) > 1 {
		username = os.Args[1]
	}

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

		go func() {
			for {
				// receive response from server
				buf := make([]byte, 64)
				_, err = conn.Read(buf)
				if err != nil {
					fmt.Printf("cannot read response: %s", err)
					os.Exit(1)
				}

				fmt.Println(string(buf))
			}
		}()

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannot read what you input\n")
		}
		fmt.Printf("\n")

		trimString := strings.TrimSpace(text)
		if trimString == "exit" {
			fmt.Println("connection closed")
			os.Exit(0)
			return
		}

		conn.Write([]byte(username + ":  " + text))
	}
}
