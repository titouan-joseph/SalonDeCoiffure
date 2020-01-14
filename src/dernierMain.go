package main

import (
	"./client"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	//"io"

	"./coiffeur"
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
func temps_process(new_client *client.Client, new_haid *coiffeur.Coiffeur) float32 {
	workingTime := 0.0
	if new_client.Sexe == "homme" {
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
func remove(s []coiffeur.Coiffeur, i int) []coiffeur.Coiffeur {   // Besoin?
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// ------ Fonction servant à placer le coiffeur dans les bonnes listes pour  la réalisation de la coupe
//		   sélectionne celui qui s'occupe du client -----

func haird_busy( coiffeursLibres chan coiffeur.Coiffeur, coiffeursOccupes chan coiffeur.Coiffeur ) coiffeur.Coiffeur {

	coiff_occupe := <- coiffeursLibres
	coiffeursOccupes <- coiff_occupe  // ajout du coiffeur dans la liste des coiffeurs occupés
	coiff_occupe.Libre = false
	return coiff_occupe
}

// ------ Fonction servant à placer le coiffeur dans les bonnes listes après la réalisation de la coupe -----

func haird_not_busy(new_haird coiffeur.Coiffeur, coiffeursLibres chan coiffeur.Coiffeur, coiffeursOccupes chan coiffeur.Coiffeur) {

	// Ecriture dans le fichier texte du client et des caractéristiques
	//EcritureClient(new_client, new_haird)

	for i := 0; i < len(coiffeursOccupes) ; i++ {

		coiff_retire_channel := <- coiffeursOccupes
		if coiff_retire_channel == new_haird {
			new_haird.Libre = true
			coiffeursLibres <- new_haird
		} else {
			coiffeursOccupes <- coiff_retire_channel
		}
	}

}

//  ----- Fonction servant à terminer la simulation -----

func end_of_day() time.Duration {

	endTimer := time.Now()                      // arret du timer
	timeOfExecution := endTimer.Sub(startTimer) // calcul du temps
	// fermer ecriture du fichier et imprime le fichier
	return timeOfExecution
}

func deleteFile( path string) {
	// delete file
	var err = os.Remove(path)
	if err != nil {
		{
			// handle error
			fmt.Printf("Could not create listener\n")
			panic(err)
		}
	}
	fmt.Println("File Deleted")

}

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			// handle error
			fmt.Printf("Could not create listener\n")
			panic(err)
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func operation(new_client *client.Client, new_haird *coiffeur.Coiffeur, coiffeursLibres chan coiffeur.Coiffeur, coiffeursOccupes chan coiffeur.Coiffeur) { // gérer file
	duration := time.Duration(temps_process(new_client, new_haird))
	EcritureClient(new_client, new_haird)
	fmt.Println(new_haird, "  prend en charge  ", new_client, " en temps: ", duration)
	time.Sleep(duration*time.Second) // effectue un équivalent de time.sleep sur la goroutine
	wg.Done()
	haird_not_busy(*new_haird, coiffeursLibres, coiffeursOccupes)
	fmt.Println(new_haird, "  a fini avec  ", new_client)

}

// ----- Fonction Main du projet -----
func test() {

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
	listeClients := CreationClients()
	fmt.Println("Creation liste clients ")

	for i:=0; i < nombreClients; i++ {
		fileAttente <-listeClients[i]
		wg.Add(1)
	}
	fmt.Println(" Coiffeurs : ", coiffeursLibres)
	fmt.Println(" Clients : ", len(fileAttente))

	for i:=0; i < nombreCoiffeurs; i++ {
		fileCoiffeursLibres <- coiffeursLibres[i]
	}


	for len(fileAttente)!= 0  { //equivalent du while qui tourne pendant toute l'execution du programme

		clientOccupe := <-fileAttente                       // retire un client de la file d'attente
		newHaird := haird_busy(fileCoiffeursLibres, fileCoiffeursOccupes)                            // choisit quel coiffeur s'en occupe
		go operation(&clientOccupe, &newHaird, fileCoiffeursLibres, fileCoiffeursOccupes)

	}

	//fmt.Println("coiffeurs :", coiffeurs)
	fmt.Println("nombre coiffeurs libres :", len(fileCoiffeursLibres))

	wg.Wait() //empêche le programme de terminer avant les go-routines

	duration :=end_of_day()
	fmt.Println( " The duration of today's process for the", nombreClients, "clients was ", duration)
	}


