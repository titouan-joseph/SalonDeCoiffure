package main

import (
	"client"
	//"salon"
	"fmt"
)

var temps_coupe_femme int //va valoir 10
var temps_coupe_homme int //va valoir 6

// fonction gérant l'arrivée d'un client dans le salon
// func client_arrival(new_client client.Client, salon salon.Salon) {
// 	size_waiting_line := salon.Waiting_line_capacity
// 	salon.Wg.Add(1) //Ajout d'un client à la liste d'attente
// }

// func haird_busy(new_client client.Client) { // effectuer un time.sleep sur la goroutine du coiffeur
// 	sexe := new_client.Sexe
// }

func main() {
	//création de la slice de coiffeurs d'après InputFile.txt
	coiffeurs := CreationCoiffeurs()
	fmt.Println(coiffeurs)
	//création de la file d'attente de clients
	fileAttente := make(chan client.Client, 10)
	client1 := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
	fileAttente <- client1
	elt := <-fileAttente
	fmt.Println("File d'attente :", elt)

	//exemple de traitment d'un client par un coiffeur par une fonction test
	coiffeurs[0].ChangeSexe(&client1)
	fmt.Println("après l'opération de", coiffeurs[0].Name, ":", client1)
}
