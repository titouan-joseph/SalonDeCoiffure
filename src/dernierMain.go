package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"client"
	"coiffeur"
	"salon"
)

var wg sync.WaitGroup

var tempsCoupeFemme float64 = 10 //va valoir 10
var tempsCoupeHomme float64 = 6  //va valoir 6
var tempsShampoo float64 = 15

var coiffeursLibres []coiffeur.Coiffeur
var coiffeursOccupes []coiffeur.Coiffeur

var client_occupe client.Client
var coiff_occupe coiffeur.Coiffeur

var startTimer = time.Now()

// ----- Fonction gérant l'arrivée d'un client dans le salon -----
func client_arrival(nouveauClient client.Client, fileAttente chan client.Client) {
	//ajout du client à la file d'attente
	fileAttente <- nouveauClient
}

// ---- Fonction qui calcule le temps du traitement du client en fonction des parametres du client et du coiffeur
func temps_process(new_client client.Client, new_haid coiffeur.Coiffeur) float64 {
	workingTime := 0.0
	if new_client.Sexe == "h" {
		workingTime = new_haid.StatCoupeHomme * tempsCoupeHomme
	} else {
		workingTime = new_haid.StatCoupeFemme * tempsCoupeFemme
	}

	if new_client.Shampoo {
		workingTime += rand.Float64() * tempsShampoo
	}

	return workingTime
}

//----- Fonction qui retire un élément d'une liste ---
func remove(s []coiffeur.Coiffeur, i int) []coiffeur.Coiffeur {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// ------ Fonction servant à modéliser l'attente par la réalisation de la coupe -----
func haird_busy(new_client client.Client, new_haird coiffeur.Coiffeur, fileAttente chan client.Client) {

	if len(coiffeursLibres) == 0 {
		// On fait quoi si y a personne de dispo pour coiffeur ?
	}
	if len(coiffeursLibres) != 0 {

		client_occupe = <-fileAttente // retire un client de la file d'attente

		coiff_occupe = coiffeursLibres[0]
		remove(coiffeursLibres, 0) // retire le premier coiffeur de la liste des coiffeurs libres

		coiffeursOccupes = append(coiffeursOccupes, coiff_occupe) // ajout du coiffeur dans la liste des coiffeurs occupés
		coiff_occupe.Libre = false

		temps_process(new_client, new_haird)

		var duration = float32(temps_process(new_client, new_haird))
		//time.sleep sur le coiffeur ( uniquement sa goroutine avec time.After(time.Duration(duration))  )
	}

}

// ----- Fonction servant en fin de coupe d'un client par le coiffeur -----
func hair_end(clientRavi client.Client, coiffeurFini coiffeur.Coiffeur) {

	// Ecriture dans le fichier texte du client et des caractéristiques
	EcritureClient(clientRavi, coiffeurFini)

	//gestion du coiffeur
	coiffeurFini.Libre = true
	coiffeursLibres = append(coiffeursLibres, coiffeurFini) // ajout du coiffeur dans la liste des coiffeurs libres
}

//  ----- Fonction servant à terminer la simulation -----
func end_of_day(sal salon.Salon) {

	// arret du timer
	endTimer := time.Now()
	timeOfExecution := endTimer.Sub(startTimer) // calcul du temps
	// fermer ecriture du fichier et imprime le fichier
}

// ----- Fonction Main du projet -----
func main() {

	coiffeurs := CreationCoiffeurs() //création de la liste de coiffeurs d'après InputFile.txt
	coiffeursLibres = coiffeurs

	fmt.Println("coiffeurs :", coiffeurs)
	fmt.Println("coiffeurs libres :", coiffeursLibres)

	fileAttente := make(chan client.Client, 10) //création de la file d'attente de clients

	wg.Wait() //empêche le programme de terminer avant les go-routines
}

//TEST
//client1 := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
//fileAttente <- client1
//elt := <-fileAttente
//fmt.Println("File d'attente :", elt)
//EcritureClient(client1, coiffeurs[0])
//coiffeurs[0].ChangeSexe(&client1)
//fmt.Println("après l'opération de", coiffeurs[0].Name, ":", client1)
//}
