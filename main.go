package main

import (
	"firstGolang/api"
	"firstGolang/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	Routes := mux.NewRouter()
	Routes.HandleFunc("/", api.HomeHandler)

	go database.PostgresConnection()
	time.Sleep(5 * time.Second)

	var Port = "3000"
	fmt.Printf("Server listening in port : %v", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", Port), Routes)
	if err != nil {
		fmt.Sprintln("Error initializing the server")
		return
	}
}
