package main

import (
	 "fmt"
	 "sync"
	 "./salon"
	 "./client"
 )


 temps_coupe_femme := 10
 temps_coupe_homme := 6
// fonction gérant l'arrivée d'un client dans le salon
func client_arrival( new_client client.Client, salon salon.Salon){
   size_waiting_line := salon.Waiting_line_capacity
   salon.Wg.Add( ) //Ajout d'un client à la liste d'attente
}

func haird_busy(new_client client.Client ){  // effectuer un time.sleep sur la goroutine du coiffeur
    sexe := new_client.Sexe

}
