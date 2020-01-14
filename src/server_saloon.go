package main

import (
	"./client"
	"./coiffeur"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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

	if _, err := os.Stat("OutputFile.txt"); err == nil {
		// path/to/whatever exists
		deleteFile("OutputFile.txt")

	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		fmt.Print(" No file names as OutputFile.txt for the moment")

	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	createFile("OutputFile.txt")

	nombreClients := 8 // Simulation à n clients
	nombreCoiffeurs := 4 // Simulation à 4 coiffeurs, attention prendre loe meme nombre que dans le fichier texte

	fileAttente := make(chan client.Client, nombreClients) //création de la file d'attente de clients
	fileCoiffeursLibres := make(chan coiffeur.Coiffeur, nombreCoiffeurs)
	fileCoiffeursOccupes := make(chan coiffeur.Coiffeur, nombreCoiffeurs)

	fmt.Println("Creation file d'attente ")
	coiffeursLibres := CreationCoiffeurs()          //création de la liste de coiffeurs d'après InputFile.txt
	fmt.Println("Creation liste coiffeurs ")

	fmt.Println(" Coiffeurs : ", coiffeursLibres)
	fmt.Println(" Clients : ", len(fileAttente))


	for {
		conn, err := server.Accept()
		if err != nil {
			// handle error
			fmt.Printf("Erreur lors de la prise en compte d'un client \n")
			panic(err)
		}

		go connTraitement(conn, fileAttente)
		go salon(fileAttente, fileCoiffeursLibres, fileCoiffeursOccupes, nombreClients)
	}

}

func salon(fileAttente chan client.Client, fileCoiffeursLibres chan coiffeur.Coiffeur,fileCoiffeursOccupes chan coiffeur.Coiffeur, nombreClients int ){

	for len(fileAttente)!= 0  { //equivalent du while qui tourne pendant toute l'execution du programme

		clientOccupe := <-fileAttente                       // retire un client de la file d'attente
		newHaird := haird_busy(fileCoiffeursLibres, fileCoiffeursOccupes)                            // choisit quel coiffeur s'en occupe
		go operation(&clientOccupe, &newHaird, fileCoiffeursLibres, fileCoiffeursOccupes)

	}

	fmt.Println("coiffeurs occupes :", len(coiffeursOccupes))

	wg.Wait() //empêche le programme de terminer avant les go-routines
	duration :=end_of_day()
	fmt.Println( " The duration of today's process for the", nombreClients, "clients was ", duration)
}


func connTraitement(connection net.Conn, file chan client.Client){

	fmt.Println("debut du traitement")

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

		file <- *tmpstruct
		wg.Add(1)
		//fmt.Println(file)


		var bin_buf bytes.Buffer
		gobStr := gob.NewEncoder(&bin_buf)
		msg := string("-> ok, client pris en charge")
		gobStr.Encode(msg)

		connection.Write(bin_buf.Bytes())
	}
}