package main

import (
	"log"
	"os"

	"./client"
	"./coiffeur"
)

// Ecrit un joli message de présentation dans OutputFile.md
func PresentationJolie() {
	fichier, err := os.OpenFile("OutputFile.md", os.O_WRONLY|os.O_APPEND, 0644)
	//constitution du message
	var message string
	message = "# SALON DE COIFFURE \n \n"
	message += "## COMPTE-RENDU DE LA JOURNEE \n"

	len, err := fichier.WriteString(message + "\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err, len)
	}
}

// Met en forme et écrit l'opération effectuée dans OutputFile.md
func EcritureClient(personne *client.Client, prestataire *coiffeur.Coiffeur) {
	fichier, err := os.OpenFile("OutputFile.md", os.O_WRONLY|os.O_APPEND, 0644)

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
	len, err := fichier.WriteString(message + "\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err, len)
	}
}

func FinJolie() {
	fichier, err := os.OpenFile("OutputFile.md", os.O_WRONLY|os.O_APPEND, 0644)
	message := "## FIN DE LA JOURNEE \n"
	len, err := fichier.WriteString("\n" + message)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err, len)
	}
}
