package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)


func main() {
	port := 666
	fmt.Printf("Ouverture serveur salon sur le  port %d\n", port)

	// the dial function connects to a server
	// the listen function creates a server

	portString := fmt.Sprintf(":%s", strconv.Itoa(port))

	ln, err := net.Listen("tcp",portString) // creation of a server on port 666 after conversion 666 to a String
	if err != nil {
		// handle error
		fmt.Printf("Could not create listener\n")
		panic(err)
	}

	connum := 1

	//using a valid listener
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Printf("Erreur lors de la prise en compte d'un client \n")
			panic(err)

		}
		go goForConnection(conn, connum) // connecting
	}
}
	func goForConnection(connection net.Conn, connum int) {

		defer connection.Close()
		connReader := bufio.NewReader(connection)

		for {
			inputLine, err := connReader.ReadString('\n')  // tries to return a single line, not including the end-of-line bytes
			if err != nil {
				fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
				fmt.Printf("Error :|%s|\n", err.Error())
				break
			}

			inputLine = strings.TrimSuffix(inputLine, "\n")  //TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged
			fmt.Printf("#DEBUG %d RCV |%s|\n", connum, inputLine)
			splitLine := strings.Split(inputLine, " ") // Splits the string every time a space is found
			returnedString := splitLine[len(splitLine)-1]
			fmt.Printf("#DEBUG %d RCV Returned value |%s|\n", connum, returnedString)
			_, _ = io.WriteString(connection, fmt.Sprintf("%s\n", returnedString))  //WriteString appends the contents of returnedString to the connection buffer. It returns the length of  returnedString and a nil error ( handled here with _, _ )
		}

	}
