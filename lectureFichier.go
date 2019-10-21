package main

import (
        "fmt"
        "io/ioutil"
		"strings"
		"strconv"
)

type coiffeur struct{
	name string
	statCoupeHomme float64
	statCoupeFemme float64
}

func main() {
        donnees, erreur := ioutil.ReadFile("InputFile.txt")
        lignes := strings.Split(string(donnees), "\n")
		var attributs []string
		var coiffeurs []coiffeur

		for i:=0; i < len(lignes); i++ {
			attributs = strings.Split(lignes[i], ":")
			prenom:= attributs[0]
			statH, err := strconv.ParseFloat(attributs[1], 32)
			statF, err := strconv.ParseFloat(attributs[2], 32)
			coiffeurs = append(coiffeurs, coiffeur{name : prenom, statCoupeHomme : statH, statCoupeFemme : statF})
			
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println(coiffeurs)
		
		if erreur != nil {
                fmt.Println(erreur)
		}
		
}