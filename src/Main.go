package main

import (
	"client"
	"coiffeur"
	"fmt"
	"salon"
	"math/rand"
	"time"
)

var tempsCoupeFemme float64 = 10 //va valoir 10
var tempsCoupeHomme float64 = 6  //va valoir 6
var tempsShampoo float64 = 15
// creation d'une liste de coiffeurs libres
var coiffeurs_libres []coiffeur.Coiffeur
// démarrage timer
var startTimer = time.Now() //Je l'ai mis la pck sinon on ne peut pas y acceder dand end_of_day

// ----- Fonction gérant l'arrivée d'un client dans le salon -----
func client_arrival(new_client client.Client, sal salon.Salon) {

	//ajout du client à la file d'attente
	sal.Wg.Add(1)
}

// ---- Fonction servant à calculer le temps que durera qui sera prit au coiffeur en fonction des parametres du client et du coiffeur
func temps_process(new_client client.Client, new_haid coiffeur.Coiffeur) float64 {
	workingTime := 0.0
	if new_client.Sexe == "h"{
		workingTime = new_haid.StatCoupeHomme * tempsCoupeHomme
	}else {
		workingTime = new_haid.StatCoupeFemme * tempsCoupeFemme
	}

	if new_client.Shampoo{
		workingTime += rand.Float64() * tempsShampoo
	}

	return workingTime
}

// ------ Fonction servant à modéliser l'attente par la réalisation de la coupe -----
func haird_busy(new_client client.Client, new_haird coiffeur.Coiffeur) {

	// retire un client de la file d'attente
	// retire un coiffeur de la liste des coiffeurs libres
	// coiffeur plus libre ( attribut )
	new_haird.Libre = false
	// appel de func temps_process
	temps_process(new_client, new_haird)
	// effectuer un time.sleep sur la goroutine du coiffeur

}

// ----- Fonction servant en fin de coupe d'un client par le coiffeur -----
func hair_end(custom client.Client, haird coiffeur.Coiffeur) {

	// Ecriture dans le fichier texte du client et des caractéristiques
	// coiffeur libre ( attribut)
	haird.Libre = true
	// ajout du coiffeur dans la liste des coiffeurs libres
	coiffeurs_libres = append(coiffeurs_libres, haird)
}

//  ----- Fonction servant à terminer la simunation -----
func end_of_day(sal salon.Salon) {

	// arret du timer
	endTimer := time.Now()
	// calcul du temps
	timeOfExecution := endTimer.Sub(startTimer)
	// fermer ecriture du fichier et imprime le fichier
}

// ----- Fonction Main du projet -----

func main() {

	//création de la liste de coiffeurs d'après InputFile.txt
	coiffeurs := CreationCoiffeurs()

	coiffeurs_libres = coiffeurs

	fmt.Println("coiffeurs :", coiffeurs)
	fmt.Println("coiffeurs libres :", coiffeurs_libres)
	//création de la file d'attente de clients
	fileAttente := make(chan client.Client, 10)
	client1 := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
	fileAttente <- client1
	elt := <-fileAttente
	fmt.Println("File d'attente :", elt)

	//test
	EcritureClient(client1, coiffeurs[0])

	//exemple de traitment d'un client par un coiffeur par une fonction test
	coiffeurs[0].ChangeSexe(&client1)
	fmt.Println("après l'opération de", coiffeurs[0].Name, ":", client1)
}
