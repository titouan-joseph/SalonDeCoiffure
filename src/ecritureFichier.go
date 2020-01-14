package main

import (
	"./coiffeur"
	client2 "coiffeur/client"
	"log"
	"os"
)

func EcritureClient(personne *client2.Client, prestataire *coiffeur.Coiffeur) {
	fichier, err := os.OpenFile("OutputFile.txt", os.O_WRONLY|os.O_APPEND, 0644)

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
	//output := []byte(message)
	len, _ := fichier.WriteString(message + "\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err, len)
	}
	//écriture du message dans OutputFile.txt
	//err := ioutil.WriteFile("OutputFile.txt", output, 0644)
	//if err != nil {
	//	print("erreur lors de l'écriture sur le fichier OutputFile")
	//}
}
