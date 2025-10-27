package crane


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

type ViewHtml struct {
	GridD [6][7]int
	TurnP int
	Joueur01 Joueur
	Joueur02 Joueur
	GGbis bool
	Tour int
	LastPlayedHTMl int
}

var ViewSite = ViewHtml{Grid, 1, Joueur1, Joueur2, GG, 0, 0}

func PlacerPièce(col int) {
	ColonPlayed = col
	Layer = 5
	for Layer >= 0 {
		if Grid[Layer][col] == 0 {
			Grid[Layer][col] = CurrentPlayer[Turn].CouleurValue
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
		
	}
	Victoire()
	ViewSite = ViewHtml{Grid, Turn, Joueur1, Joueur2, GG, +1, 0}		
}


func Reset() {
	for i := range Grid {
		for j := range Grid[i] {
			Grid[i][j] = 0
		}
	}
	for i := range ViewSite.GridD {
		for j := range ViewSite.GridD[i] {
			ViewSite.GridD[i][j] = 0
		}
	}
	GG = false
	Turn = 0
	ViewSite = ViewHtml{Grid, Turn, Joueur1, Joueur2, GG, 0, 0} 
}




