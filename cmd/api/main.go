package main

import (
	"fmt"
	"os"
	"strconv"

	"botgauge-app/internal/server"
)

func main() {
	server := server.NewServer()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	fmt.Printf("Server starting on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
