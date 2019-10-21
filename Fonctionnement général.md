Le projet est une simulation en temps maîtrisé du fonctionnement d'un salon de coiffure.

Un premier type de données nommé coiffeur comprendra un nom, ses statistiques pour les coiffures hommes, ses statistiques pour les coupes femmes
Un second type de données nommé client comprendra le sexe de celui- ci ( et donc le type de coupe attendue), et si il demande à avoir un shampoing.
Un dernier type de données nommé salon est décrit par son nom, son nombre de coiffeurs ( un input saisi par l'utilisateur), et sa capacité de file d'attente

Il n'existe qu'une coupe homme, et une coupe femme.
La simulation démarre avec la lecture d'un fichier texte pour la descritpion du salon. L'utilisateur choisit le nombre de coiffeurs présents dans le salon.
Un coiffeur occupé sera modélisé par une fonction time.sleep()

Une goroutine sera générée pour chaque coiffeur. Une autre goroutine sera initialisée dès le début de la simulation pour gérer l'arrivée aléatoire de clients.
Si un coiffeur est disponible à l'arrivée du client, il le prend en charge. Sinon, il est placé dans la file d'attente. 
La file d'attente a une capacité max d'occupation. Une fois pleine, les nouveaux clients ne peuvent plus attendre pour se faire coiffer.

A chaque client prit en charge, on écrit dans un fichier texte la descritpion du client, le coiffeur le prenant en charge, et le temps de prise en charge, pour avoir un registre complet en fin de journée.