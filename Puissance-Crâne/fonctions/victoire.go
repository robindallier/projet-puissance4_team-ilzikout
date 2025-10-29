package crane

import (
	"fmt"
	"time"
)

var VerifIndex int
var Won int
var LayerPlayed int
var GG bool = false
var LayerForDiag int
var DrawIndex int

func Victoire() {
	Won = 0
	VerifIndex = 6
	if DrawIndex == 42 && Won == 0 {
		fmt.Print("Match nul")
	}
	for VerifIndex >= 0 {
		if Won == 4 {
			GG = true
			return
		}
		if Grid[LayerPlayed][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won++
		} else {
			Won = 0
		}
		VerifIndex--
	}

	if Won >= 4 {
		GG = true
		return
	}
	Won = 0

	VerifIndex = 5
	for VerifIndex >= 0 {
		if Won == 4 {
			GG = true
			return
		}
		if Grid[VerifIndex][ColonPlayed] == CurrentPlayer[LastPlayed].CouleurValue {
			Won++
		} else {
			Won = 0
		}
		VerifIndex--
	}

	if Won >= 4 {
		GG = true
		return
	}
	Won = 0
	VerifIndex = ColonPlayed
	LayerForDiag = LayerPlayed

	for VerifIndex < 6 && LayerForDiag < 5 {
		VerifIndex++
		LayerForDiag++
	}

	for VerifIndex >= 0 && LayerForDiag >= 0 {
		if LayerForDiag < 0 || LayerForDiag >= 6 || VerifIndex < 0 || VerifIndex >= 7 {
			break
		}
		if Grid[LayerForDiag][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won++
		} else {
			Won = 0
		}
		if Won >= 4 {
			GG = true
			return
		}
		VerifIndex--
		LayerForDiag--
	}

	Won = 0
	VerifIndex = ColonPlayed
	LayerForDiag = LayerPlayed

	for VerifIndex-1 >= 0 && LayerForDiag+1 <= 5 {
		VerifIndex--
		LayerForDiag++
		if VerifIndex == -1 {
			VerifIndex = 0
			break
		}
	}

	for VerifIndex <= 6 && LayerForDiag >= 0 {
		if Grid[LayerForDiag][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won++
		} else {
			Won = 0
		}
		if Won >= 4 {
			GG = true
			return
		}
		VerifIndex++
		LayerForDiag--
	}
}

func WinMessage(S string) {
	for _, Letter := range S {
		fmt.Printf("%c", Letter)
		time.Sleep(50 * time.Millisecond)
	}
}
