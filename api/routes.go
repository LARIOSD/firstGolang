package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Printf("\nStart Counter")
	for i := 0; i < 10; i++ {
		fmt.Printf("\nCounter : %d ", i)
	}
	// time.Sleep(3 * time.Minute)
	writer.Write([]byte("Hello word !"))
}
