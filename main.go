package main

import (
	"fmt"
	"net/http"
	"os"

	"./app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthemtication) //attach JWT auth middleware

	port := os.Getenv("PORT") // Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the ap, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}