package crane
import "fmt"
func Victoire(grille [][]int) bool {
    lignes := len(grille)
    colonnes := len(grille[0])

    for i := 0; i < lignes; i++ {
        for j := 0; j < colonnes; j++ {
            joueur := grille[i][j]
            if joueur == 0 {
                continue
            }
            if j+3 < colonnes {
                horizontal := true
                for k := 0; k < 4; k++ {
                    if grille[i][j+k] != joueur {
                        horizontal = false
                        break
                    }
                }
                if horizontal {
                    return true
                }
            }
            if i+3 < lignes {
                vertical := true
                for k := 0; k < 4; k++ {
                    if grille[i+k][j] != joueur {
                        vertical = false
                        break
                    }
                }
                if vertical {
                    return true
                }
            }
            if i+3 < lignes && j+3 < colonnes {
                diagDesc := true
                for k := 0; k < 4; k++ {
                    if grille[i+k][j+k] != joueur {
                        diagDesc = false
                        break
                    }
                }
                if diagDesc {
                    return true
                }
            }
            if i+3 < lignes && j-3 >= 0 {
                diagAsc := true
                for k := 0; k < 4; k++ {
                    if grille[i+k][j-k] != joueur {
                        diagAsc = false
                        break
                    }
                }
                if diagAsc {
                    return true
                }
            }
        }
    }

    return false
}

// test

var TestG1 = [][]int{
    {0, 0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0, 0},
    {0, 0, 1, 2, 0, 0, 0},
    {0, 1, 2, 2, 0, 0, 0},
    {1, 2, 1, 1, 0, 0, 0},
    {2, 1, 2, 1, 0, 0, 0},
}


