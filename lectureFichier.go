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
		attributs := make([]string, len(lignes))
		coiffeurs := make([]coiffeur, len(lignes))

		for i:=0; i < len(lignes); i++ {
			attributs = append(strings.Split(lignes[i], ":"))
			//conversion du type string au type float32
			prenom:= attributs[0]
			statH, err := strconv.ParseFloat(attributs[1], 32)
			statF, err := strconv.ParseFloat(attributs[2], 32)
			coiffeurs = append(coiffeurs, coiffeur{name : prenom, statCoupeHomme : statH, statCoupeFemme : statF})
			
			if err != nil {
				break
			}
		}
        if erreur != nil {
                fmt.Println(erreur)
		}
		
		fmt.Println(coiffeurs)
}