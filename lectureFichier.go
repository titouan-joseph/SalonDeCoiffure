package main

import (
        "fmt"
        "io/ioutil"
        "strings"
)

func main() {
        donnees, erreur := ioutil.ReadFile("texte.txt")
        lignes := strings.Split(string(donnees), "\n")

        if erreur != nil {
                fmt.Println(erreur)
        }

        fmt.Println(lignes[1])
}