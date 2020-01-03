package coiffeur

import (
	
	"./client"
)

type Coiffeur struct {
	Name           string
	StatCoupeHomme float64
	StatCoupeFemme float64
	Libre          bool
}

//fonction test temporaire, juste pour avoir la syntaxe
func (coiff Coiffeur) ChangeSexe(personne *client.Client) {
	if personne.Sexe == "homme" {
		personne.Sexe = "femme"
	} else if personne.Sexe == "femme" {
		personne.Sexe = "homme"
	}
}
