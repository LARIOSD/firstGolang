package main

import (
	"firstGolang/api"
	"firstGolang/config"
	"firstGolang/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Routes := mux.NewRouter()
	Routes.HandleFunc("/", api.HomeHandler)

	config.LoadEnv()
	database.PostgresConnection()

	environment := config.GetEnvironment()

	fmt.Printf("Server listening in port : %v", environment.ServerPort)

	err := http.ListenAndServe(fmt.Sprintf(":%v", environment.ServerPort), Routes)
	if err != nil {
		fmt.Sprintln("Error initializing the server")
		return
	}
}
