package coiffeur

import (
	client2 "coiffeur/client"
)

type Coiffeur struct {
	Name           string
	StatCoupeHomme float64
	StatCoupeFemme float64
	Libre          bool
}

//fonction test temporaire, juste pour avoir la syntaxe
func (coiff Coiffeur) ChangeSexe(personne *client2.Client) {
	if personne.Sexe == "homme" {
		personne.Sexe = "femme"
	} else if personne.Sexe == "femme" {
		personne.Sexe = "homme"
	}
}
