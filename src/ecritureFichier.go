package main

import (
	"client"
	"io/ioutil"
)

func EcritureClient(personne client.Client) {
	nom_client := []byte(personne.Name)
	err := ioutil.WriteFile("OutputFile.txt", nom_client, 0644)
	if err != nil {
		print("erreur lors de l'écriture sur le fichier OutputFile")
	}
}
