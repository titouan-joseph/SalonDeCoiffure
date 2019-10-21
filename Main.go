package main

import (
	 "fmt"
	 "sync"
   "salon"
   "client"
 )


 temps_coupe_femme := 10
 temps_coupe_homme := 6
// fonction gérant l'arrivée d'un client dans le salon
func client_arrival( new_client client, salon salon){
   size_waiting_line := salon.waiting_line_capacity
   salon.wg.Add( )  //Ajout d'un client à la liste d'attente
}

func haird_busy(new_client client ){  // effectuer un time.sleep sur la goroutine du coiffeur
    sexe= new_client.sexe

}
