package main

import (
	"bytes"
	"client"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// server.go


func main() {

	port := 666
	addr := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port)) //String conversion to use Dial function

	server, err := net.Listen("tcp", addr)
	if err != nil {
		// handle error
		fmt.Printf("Could not create listener\n")
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			// handle error
			fmt.Printf("Erreur lors de la prise en compte d'un client \n")
			panic(err)
		}

		go connTraitement(conn)
	}
}

func connTraitement(connection net.Conn){

	defer connection.Close()

	tmp := make([]byte, 512)

	for {

		_, err := connection.Read(tmp)
		if err != nil {
			fmt.Printf("#DEBUG RCV ERROR no panic, just a client\n")
			fmt.Printf("Error :|%s|\n", err.Error())
			break
		}
		tmpbuff := bytes.NewBuffer(tmp)
		tmpstruct := new(client.Client)
		gobobj := gob.NewDecoder(tmpbuff)
		gobobj.Decode(tmpstruct)

		fmt.Println(tmpstruct)

		var bin_buf bytes.Buffer
		gobStr := gob.NewEncoder(&bin_buf)
		msg := string("-> ok, client pris en charge")
		gobStr.Encode(msg)

		connection.Write(bin_buf.Bytes())
	}
}