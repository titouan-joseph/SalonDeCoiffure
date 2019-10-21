package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coiffeur struct {
	name           string
	statCoupeHomme float64
	statCoupeFemme float64
}

func main() {
	donnees, erreur := ioutil.ReadFile("InputFile.txt")

	if erreur != nil {
		fmt.Println("Erreur lors de la lecture du fichier")
	}

	lignes := strings.Split(string(donnees), "\n")
	var attributs []string
	var coiffeurs []coiffeur

	for i := 0; i < len(lignes); i++ {
		attributs = strings.Split(lignes[i], ":")
		prenom := attributs[0]
		statH, _ := strconv.ParseFloat(attributs[1], 64)
		statF, _ := strconv.ParseFloat(attributs[2], 64)
		coiffeurs = append(coiffeurs, coiffeur{name: prenom, statCoupeHomme: statH, statCoupeFemme: statF})
	}

	fmt.Println(coiffeurs)
}
