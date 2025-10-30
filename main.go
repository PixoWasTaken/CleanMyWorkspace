package main // Les commentaires je les mets juste pour moi meme faites pas attention

import (
	"CleanWork/Clean"

	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)


func ShowFika(fika *[][]*string) { // Pas grand chose d'intéressant, compliqué mais en gros on fait le tour de la fonction pour afficher les chaises et les items
	for _, row := range *fika {
		fmt.Print("|")
		for _, item := range row {
			if item == nil {
				fmt.Print("Chaise|") // ça fait plus de sens que c'était des chaises, plutot que "nil"
			}else {
				fmt.Printf("%s|", *item)
			}
		}
	}
}

func main() {
	fika := CleanMyWorkspace.GenererateWorkSpace()

	fmt.Println("Fika sale :")
	ShowFika(fika)

	fika = Clean.CleanWorkSpace(fika) // Une fois que la fika est affichée avec les items, on change au millieu, en appelant la fonction de la fika propre

	fmt.Println("\n ")
	fmt.Println("\nFika nettoyée")
	ShowFika(fika)
}