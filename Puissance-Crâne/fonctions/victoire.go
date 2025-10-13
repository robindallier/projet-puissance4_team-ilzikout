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

func Victoire() {
	Won = 0
	VerifIndex = 6

	for VerifIndex >= 0 {
		if Won == 4 {
			GG = true
			fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
			return
		}
		if Grid[LayerPlayed][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won ++
		} else {
			Won = 0
		}
		VerifIndex -- 
	}

	if Won >= 4 {
		GG = true
		fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
		return
	} 
	Won = 0

	VerifIndex = 5
	for VerifIndex >= 0 {
		if Won == 4 {
			GG = true
			fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
			return
		}
		if Grid[VerifIndex][ColonPlayed-1] == CurrentPlayer[LastPlayed].CouleurValue {
			Won ++
		} else {
			Won = 0
		}
		VerifIndex --
	}

	if Won >= 4 {
		GG = true
		fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
		return
	} 
	Won = 0
	VerifIndex = ColonPlayed -1
	LayerForDiag = LayerPlayed

	for VerifIndex < 6 && LayerForDiag < 5 {
			VerifIndex ++
			LayerForDiag ++
	}

	for VerifIndex >= 0 && LayerForDiag >= 0 {	
		if Grid[LayerForDiag][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won ++
		} else {
			Won = 0
		}
		if Won >= 4 {
			GG = true
			fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
			return
		}
		VerifIndex --
		LayerForDiag --
	}

	Won = 0
	VerifIndex = ColonPlayed -1
	LayerForDiag = LayerPlayed

	for VerifIndex - 1 >= 0 && LayerForDiag + 1 <= 5 {
			VerifIndex --
			LayerForDiag ++
	}

	for VerifIndex <= 6 && LayerForDiag >= 0 {	
		if Grid[LayerForDiag][VerifIndex] == CurrentPlayer[LastPlayed].CouleurValue {
			Won ++
		} else {
			Won = 0
		}
		if Won >= 4 {
			GG = true
			fmt.Print(CurrentPlayer[LastPlayed].Name);WinMessage(" a gagné\n",)
			return
		}
		VerifIndex ++
		LayerForDiag --
	}
}

func WinMessage(S string) {
	for _, Letter := range S {
		fmt.Printf("%c", Letter)
		time.Sleep(50 * time.Millisecond)
	}
}
