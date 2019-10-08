# SalonDeCoiffure
Modélisation d'un salon de coiffure en Go.


L'idée est de modélisé la file d'attente d'un salon de coiffure pour optimiser l'attente des clients.
Les coiffeurs et les clients seront modélisé par des thread
Un salon est composé de sièges pour l'attente des clients
Si tous les sièges sont occupés, les clients repartent
Chaque client a un genre, le coiffeur met plus de temps a coiffer une femme qu'un homme
A leurs arrivé les clients savent s'ils doivent se faire shampouiner et leurs préférences de coupe aux ciseaux ou a la tondeuse
Ils s'asseyent et attendent leur tour


Un coiffeur met plus au moins de temps à shampouiner un client dans un interval de temps
Il met plus de temps à faire une coupe au ciseaux qu'a la tondeuse
