package main

import (
	"go_modules/routes"
	"net/http"

	_ "github.com/lib/pq"
)


func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

