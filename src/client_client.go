package main

import (
	"bytes"
	"client"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// client.go


func main(){

	port := 666
	addr := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))  //String conversion to use Dial function

	conn, err := net.Dial("tcp", addr)

	if err != nil{
		//handle error
		fmt.Printf("Connection impossible au salon sur le port  %d\n", port)
		os.Exit(1)
	} else {

		defer conn.Close()

		fabrice := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
		sophie := client.Client{Name: "Sophie", Sexe: "femme", Shampoo: true}
		thomas := client.Client{Name: "Thomas", Sexe: "homme", Shampoo: true}
		thomas1 := client.Client{Name: "Thomas1", Sexe: "homme", Shampoo: true}
		thomas2 := client.Client{Name: "Thomas2", Sexe: "homme", Shampoo: true}
		thomas3 := client.Client{Name: "Thomas3", Sexe: "homme", Shampoo: true}
		thomas4 := client.Client{Name: "Thomas4", Sexe: "homme", Shampoo: true}
		thomas5 := client.Client{Name: "Thomas5", Sexe: "homme", Shampoo: true}

		clientList := []client.Client{fabrice, sophie, thomas, thomas1, thomas2, thomas3, thomas4, thomas5}

		// encode buffer and marshal it into a gob object
		tmp := make([]byte, 128)

		for _, clt := range clientList {
			fmt.Print(clt)
			time.Sleep(1 * time.Second)

			var bin_buf bytes.Buffer
			gobobj := gob.NewEncoder(&bin_buf)
			gobobj.Encode(clt)

			conn.Write(bin_buf.Bytes())

			_, _ = conn.Read(tmp)
			tmpbuff := bytes.NewBuffer(tmp)
			tmpStr := new(string)
			gobStr := gob.NewDecoder(tmpbuff)
			gobStr.Decode(tmpStr)

			fmt.Println("recu by server:", tmpStr)
		}
	}
}