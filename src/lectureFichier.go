package main

import (
	"./coiffeur"
	client2 "coiffeur/client"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func CreationCoiffeurs() []coiffeur.Coiffeur {
	donnees, erreur := ioutil.ReadFile("InputFileCoiffeurs.txt")

	if erreur != nil {
		fmt.Println("Erreur lors de la lecture du fichier")
	}

	lignes := strings.Split(string(donnees), "\n")
	var attributs []string
	var Coiffeurs []coiffeur.Coiffeur

	for i := 0; i < len(lignes); i++ {
		attributs = strings.Split(lignes[i], ":")
		prenom := attributs[0]
		statH, _ := strconv.ParseFloat(attributs[1], 64)
		statF, _ := strconv.ParseFloat(attributs[2], 64)
		Coiffeurs = append(Coiffeurs, coiffeur.Coiffeur{Name: prenom, StatCoupeHomme: statH, StatCoupeFemme: statF})
	}

	return Coiffeurs
}

func CreationClients() []client2.Client {
	donnees, erreur := ioutil.ReadFile("InputFileClients.txt")

	if erreur != nil {
		fmt.Println("Erreur lors de la lecture du fichier")
	}

	lignes := strings.Split(string(donnees), "\n")
	var attributs []string
	var Clients []client2.Client

	for i := 0; i < len(lignes); i++ {
		attributs = strings.Split(lignes[i], ":")
		prenom := attributs[0]
		sexe := attributs[1]
		shampStr:= attributs[2]
		var shampBool bool
		if shampStr == "true" {
			shampBool = true
		} else {
			shampBool = false
		}


		Clients = append(Clients, client2.Client{Name: prenom, Sexe: sexe, Shampoo: shampBool})
	}

	return Clients
}