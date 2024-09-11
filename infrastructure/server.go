package infrastructure

import (
	"fmt"
	"net/http"
)

func StartHTTPServer(port string) {
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
