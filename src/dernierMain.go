package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"./client"
	"./coiffeur"
	//"./salon"
)

var wg sync.WaitGroup

var tempsCoupeFemme float64 = 10 //va valoir 10
var tempsCoupeHomme float64 = 6  //va valoir 6
var tempsShampoo float64 = 150000

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
func temps_process(new_client *client.Client, new_haid *coiffeur.Coiffeur) float32 {
	workingTime := 0.0
	if new_client.Sexe == "h" {
		workingTime = new_haid.StatCoupeHomme * tempsCoupeHomme
	} else {
		workingTime = new_haid.StatCoupeFemme * tempsCoupeFemme
	}

	if new_client.Shampoo {
		workingTime += rand.Float64() * tempsShampoo
	}
	return float32(workingTime)
}

//----- Fonction qui retire un élément d'une liste ---
func remove(s []coiffeur.Coiffeur, i int) []coiffeur.Coiffeur {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// ------ Fonction servant à placer le coiffeur dans les bonnes listes pour  la réalisation de la coupe
//		   sélectionne celui qui s'occupe du client -----

func haird_busy() coiffeur.Coiffeur {

	coiff_occupe = coiffeursLibres[0]
	coiffeursLibres= remove(coiffeursLibres, 0)                                // retire le premier coiffeur de la liste des coiffeurs libres
	coiffeursOccupes = append(coiffeursOccupes, coiff_occupe) // ajout du coiffeur dans la liste des coiffeurs occupés
	coiff_occupe.Libre = false
	return coiff_occupe
}

// ------ Fonction servant à placer le coiffeur dans les bonnes listes après la réalisation de la coupe -----

func haird_not_busy(new_haird coiffeur.Coiffeur, new_client client.Client) {

	// Ecriture dans le fichier texte du client et des caractéristiques
	//EcritureClient(new_client, new_haird)

	coiffeursLibres = append(coiffeursLibres, new_haird)
	for i := 0; i < len(coiffeursOccupes); i++ {
		if coiffeursOccupes[i] == new_haird {
			coiffeursOccupes= remove(coiffeursOccupes, i)
		}
		coiff_occupe.Libre = true
	}

}

//  ----- Fonction servant à terminer la simulation -----

func end_of_day() float64 {

	endTimer := time.Now()                      // arret du timer
	timeOfExecution := endTimer.Sub(startTimer) // calcul du temps
	// fermer ecriture du fichier et imprime le fichier
	return float64(timeOfExecution)
}

func operation(new_client *client.Client, new_haird *coiffeur.Coiffeur, fileAttente chan client.Client) {
	duration := temps_process(new_client, new_haird)
	fmt.Println(new_haird, "  prend en charge  ", new_client, " en temps: ", duration)
	time.Sleep(10000000000) // effectue un équivalent de time.sleep sur la goroutine
	wg.Done()

}

// ----- Fonction Main du projet -----
func main() {


	nombreClients := 8 // Simulation à n clients
	fileAttente := make(chan client.Client, nombreClients) //création de la file d'attente de clients
	coiffeursZizi := CreationCoiffeurs()           //création de la liste de coiffeurs d'après InputFile.txt

	fmt.Println(" Coiffeurs : ", coiffeursZizi)

	fabrice := client.Client{Name: "Fabrice", Sexe: "homme", Shampoo: false}
	fileAttente <- fabrice
	sophie := client.Client{Name: "Sophie", Sexe: "femme", Shampoo: true}
	fileAttente <- sophie
	thomas := client.Client{Name: "Thomas", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas
	thomas1 := client.Client{Name: "Thomas1", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas1
	thomas2 := client.Client{Name: "Thomas2", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas2
	thomas3 := client.Client{Name: "Thomas3", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas3
	thomas4 := client.Client{Name: "Thomas4", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas4
	thomas5 := client.Client{Name: "Thomas5", Sexe: "homme", Shampoo: true}
	fileAttente <- thomas5

	coiffeursLibres =coiffeursZizi


	for i := 0; i < nombreClients; i++ {
		wg.Add(1) // il y aura maximum nombreClients go-routines. Les ajoute pour les préparer à être utilisées
	}

	for len(fileAttente) != 0 && len(coiffeursLibres) != 0 { //equivalent du while qui tourne pendant toute l'execution du programme

		clientOccupe := <-fileAttente                       // retire un client de la file d'attente
		newHaird := haird_busy()                            // choisit quel coiffeur s'en occupe
		go operation(&clientOccupe, &newHaird, fileAttente) //   ----  ajouter les arguments
		haird_not_busy(newHaird, clientOccupe)              // equivalent de haird_busy mais en fin de traitement ( gère qui est de nouveau dispo)
	}

	//fmt.Println("coiffeurs :", coiffeurs)
	fmt.Println("coiffeurs libres :", coiffeursLibres)
	fmt.Println("coiffeurs occupes :", coiffeursOccupes)

	wg.Wait() //empêche le programme de terminer avant les go-routines

	duration :=end_of_day()
	fmt.Println( " The duration of today's process for the", nombreClients, "clients was ", duration)
	}


