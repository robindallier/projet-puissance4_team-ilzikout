package main	

import (
	"fmt"
	"net/http"
	"power4/router"
)

func main() {
	r := router.New()

	fmt.Println("serveur démare sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
