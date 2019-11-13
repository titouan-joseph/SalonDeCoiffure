package main

import (
	"client"
	"coiffeur"
	"io/ioutil"
)

func EcritureClient(personne client.Client, prestataire coiffeur.Coiffeur) {
	//constitution du message
	var message string
	message = personne.Name
	message += " pris(e) en charge par "
	message += prestataire.Name

	if personne.Shampoo {
		message += " avec shampooing"
	} else {
		message += " sans shampooing"
	}
	output := []byte(message)

	//écriture du message dans OutputFile.txt
	err := ioutil.WriteFile("OutputFile.txt", output, 0644)
	if err != nil {
		print("erreur lors de l'écriture sur le fichier OutputFile")
	}
}
