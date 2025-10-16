package main

import (
	"fmt"
	"net/http"
	"power4/routeur"
)

func main() {
	r := routeur.New()
	
	fmt.Println("serveur démare sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
