package puissance4local

import "fmt"

type Joueur struct {
	Name string
	Couleur string
	CouleurValue int
	Victoires int
}

var Joueur1 = Joueur{"Jaune", "_", 1, 0}
var Joueur2 = Joueur{"Rouge", "_", 2, 0}

var JRouge int = 1
var JJaune int = 2

var LastPlayed int

var Grid[6][7]int

var ColonPlayed int
var Layer int = 5

var CurrentPlayer = []*Joueur{&Joueur1, &Joueur2}

var Turn int = 0

func PlacerPièce() {		
	Printgrid()
	Layer = 5
	fmt.Print("ou voulez vous placer votre jeton ?\n")
	fmt.Scan(&ColonPlayed)
	for Layer >= 0 {
		if Grid[Layer][ColonPlayed-1] == 0 {
			fmt.Print("vous avez placé votre jeton sur la ", ColonPlayed, "ème colone\n")
			Grid[Layer][ColonPlayed-1] = CurrentPlayer[Turn].CouleurValue
			LayerPlayed = Layer
			if Turn == 0 {
				Turn = 1
				LastPlayed = 0
			} else {
				Turn = 0
				LastPlayed = 1
			}
			break
		} else {
			Layer --
		}
	}
	if Layer < 0 {
		fmt.Print("vous ne pouvez pas placer sur cette colone\n")
	}			
}

func Printgrid() {
	for _, Ligne := range Grid {
		fmt.Print(Ligne, "\n")
	}
}

func TPT() {
	RecupName()
	for !GG {
		PlacerPièce()
		Victoire()
	}	
}


func RecupName() {
	fmt.Print("Donnez le nom du joueur 1\n")
	fmt.Scan(&Joueur1.Name)
	fmt.Print("Donnez le nom du joueur 2\n")
	fmt.Scan(&Joueur2.Name)
}
