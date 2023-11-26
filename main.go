package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Routes := mux.NewRouter()
	Routes.HandleFunc("/", func(writer http.ResponseWriter, reader *http.Request) {
		writer.Write([]byte("Hello word !"))
	})

	var Port = "3000"

	fmt.Printf("Server listening in port : %v", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", Port), Routes)
	if err != nil {
		fmt.Sprintln("Error initializing the server")
		return
	}
}
