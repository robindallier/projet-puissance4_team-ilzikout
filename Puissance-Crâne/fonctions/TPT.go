package crane

type Joueur struct {
	Name         string
	Couleur      string
	CouleurValue int
	Victoires    int
}

var Joueur1 = Joueur{"Jaune", "_", 1, 5}
var Joueur2 = Joueur{"Rouge", "_", 2, 0}

var JRouge int = 1
var JJaune int = 2

var LastPlayed int

var Grid [6][7]int

var ColonPlayed int
var Layer int = 5

var CurrentPlayer = []*Joueur{&Joueur1, &Joueur2}

var Turn int = 0

type ViewHtml struct {
	GridD          [6][7]int
	TurnP          int
	Joueur01       Joueur
	Joueur02       Joueur
	GGbis          bool
	IsDraw         bool
	Tour           int
	LastPlayedHTMl int
	CurrentPlayer  string
	Named			bool
}

var DrawCount int

var ViewSite = ViewHtml{Grid, 1, Joueur1, Joueur2, GG, false, 0, 0, CurrentPlayer[Turn].Name, false}

var Draw bool = false

func PlacerPièce(col int) {
	DrawCount ++
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
			Layer--
		}
	}

	if DrawCount == 42 {
		Draw = true
	} else {
		Draw = false
	}

	Victoire()
	if GG {
		CurrentPlayer[LastPlayed].Victoires ++
	}
	ViewSite = ViewHtml{Grid, Turn, Joueur1, Joueur2, GG, Draw, +1, 0, CurrentPlayer[Turn].Name, true}
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
	Joueur1.Victoires = 0
	Joueur2.Victoires = 0
	ViewSite.Named = false
	ViewSite = ViewHtml{Grid, 1, Joueur1, Joueur2, GG, false, 0, 0, CurrentPlayer[Turn].Name, false}
}

func Restart() {
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
	ViewSite = ViewHtml{Grid, Turn, Joueur1, Joueur2, GG, false, 0, 0, CurrentPlayer[Turn].Name, true}
}
