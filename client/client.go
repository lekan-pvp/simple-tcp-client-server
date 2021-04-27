package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//open connection:
	conn, err := net.Dial("tcp", ":50000")
	if err != nil {
		//no connection could be made because the target machine actively refused it
		fmt.Println("Error dialing", err.Error())
		return //terminate program
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	//send into server until Quit:
	for {
		fmt.Println("What to send to the server? Type Q to Quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n"))
	}
}
