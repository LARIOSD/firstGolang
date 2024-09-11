package main

import (
	"firstGolang/api"
	"firstGolang/environment/config"
	"firstGolang/environment/upload"
	postgres "firstGolang/postgres/connection"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Routes := mux.NewRouter()
	Routes.HandleFunc("/prueba", api.HomeHandler)

	upload.UploadEnv()

	postgres.NewConnectPostgres()
	environment := config.GetEnvironment()

	fmt.Printf("Server listening in port : %v", environment.ServerPort)

	err := http.ListenAndServe(fmt.Sprintf(":%v", environment.ServerPort), Routes)
	if err != nil {
		fmt.Sprintln("Error initializing the server")
		return
	}
}

// run project --> air
