package main

import (
	"client"
	"coiffeur"
	"fmt"
	"salon"
)

var temps_coupe_femme int //va valoir 10
var temps_coupe_homme int //va valoir 6

// ----- Fonction gérant l'arrivée d'un client dans le salon -----
func client_arrival(new_client client.Client, sal salon.Salon) {

	//ajout du client à la file d'attente

}

// ---- Fonction servant à calculer le temps que durera qui sera prit au coiffeur en fonction des parametres du client et du coiffeur
func temps_process(new_client client.Client, new_haid coiffeur.Coiffeur) {

}

// ------ Fonction servant à modéliser l'attente par la réalisation de la coupe -----
func haird_busy(new_client client.Client, new_haird coiffeur.Coiffeur) {

	// retire un client de la file d'attente
	// retire un coiffeur de la liste des coiffeurs libres
	// coiffeur plus libre ( attribut )
	// appel de func temps_process
	// effectuer un time.sleep sur la goroutine du coiffeur

}

// ----- Fonction servant en fin de coupe d'un client par le coiffeur -----
func hair_end(custom client.Client, haird coiffeur.Coiffeur) {

	// Ecriture dans le fichier texte du client et des caractéristiques
	// coiffeur libre ( attribut)
	// ajout du coiffeur dans la liste des coiffeurs libres
}

//  ----- Fonction servant à terminer la simunation -----
func end_of_day(sal salon.Salon) {

	// arret du timer
	// calcul du temps
	// fermer ecriture du fichier et imprime le fichier
}

// ----- Fonction Main du projet -----

func main() {

	//création de la liste de coiffeurs d'après InputFile.txt
	coiffeurs := CreationCoiffeurs()

	// creation d'une liste de coiffeurs libres
	var coiffeurs_libres []coiffeur.Coiffeur
	coiffeurs_libres = coiffeurs

	// démarrage timer

	fmt.Println("coiffeurs :", coiffeurs)
	fmt.Println("coiffeurs libres :", coiffeurs_libres)
	//création de la file d'attente de clients
	fileAttente := make(chan client.Client, 10)
	client1 := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
	fileAttente <- client1
	elt := <-fileAttente
	fmt.Println("File d'attente :", elt)

	//test
	EcritureClient(client1)

	//exemple de traitment d'un client par un coiffeur par une fonction test
	coiffeurs[0].ChangeSexe(&client1)
	fmt.Println("après l'opération de", coiffeurs[0].Name, ":", client1)
}
