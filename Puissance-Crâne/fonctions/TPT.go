package crane

var Name1 string
var Name2 string

var Played bool

var CurrentPlaying int

type Jeton struct {
	Couleur int
}

var JRouge = Jeton{1}
var JJaune = Jeton{2}

type Joueur struct {
	Couleur int
	name string
}

var Joueur1 = Joueur{JRouge.Couleur, Name1}
var Joueur2 = Joueur{JJaune.Couleur, Name2}
var Players = []Joueur{Joueur1, Joueur2}

var Grid[7][6]int
var Layer int

func Placed(SelectedColone int) {
	Layer = 5
	Played = false
	for Layer >= 0 || !Played{
		if Grid[SelectedColone][Layer] == 0 {
			Grid[SelectedColone][Layer] = Players[CurrentPlaying].Couleur
			Played = true
			break
		} else {
			Layer --
		}
	}
	if !Played {
		// message de pas possible
	}		
}
