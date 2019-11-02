package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"./coiffeur"
)

func CreationCoiffeurs() []coiffeur.Coiffeur {
	donnees, erreur := ioutil.ReadFile("InputFile.txt")

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
