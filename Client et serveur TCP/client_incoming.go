package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)


func main() {

	port := 666
	fmt.Printf("TCP Server du client sur le port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))  //String conversion to use Dial function

	conn, err := net.Dial("tcp",  portString)

	if err != nil {
		//handle error
		fmt.Printf("Connection impossible au salon sur le port  %d\n", port)
		os.Exit(1)
	} else {

		defer conn.Close()
		reader := bufio.NewReader(conn)
		fmt.Printf("Client connecté au salon \n")
		for i:= 0; i < 10; i++{

			io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i)) //test ecriture fichier

			resultString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Main could not read from server")
				os.Exit(1)
			}
			resultString = strings.TrimSuffix(resultString, "\n")  //TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
			fmt.Printf(" Le salon a répondu  : |%s|\n", resultString)


		}

	}
}
