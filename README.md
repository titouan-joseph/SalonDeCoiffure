# SalonDeCoiffure
Modélisation d'un salon de coiffure en Go.


L'idée est de modéliser la file d'attente d'un salon de coiffure pour réduire du mieux possible l'attente des clients.
Les coiffeurs et les clients seront modélisé par des thread et la file de clients par une pile FIFO.
Un salon est composé de sièges pour l'attente des clients
Si tous les siéges sont occupés, les clients repartent
Chaque client a un genre, le coiffeur met plus de temps a coiffer une femme qu'un homme
A leur arrivée les clients savent s'il doivent se faire shampouiner ainsi que leurs préférences de coupe aux ciseaux ou a la tondeuse.
Ils s'asseyent et attendent leur tour


Un coiffeur met plus au moins de temps à shampouiner un client.
Il met plus de temps a faire une coupe au ciseaux qu'a la tondeuse.
