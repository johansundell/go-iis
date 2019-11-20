package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Default port
	port := "8080"
	// Try to get the post from .net core
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	srv := &http.Server{
		Addr: ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	http.HandleFunc("/", mainHandler)

	log.Println(srv.ListenAndServe())
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! :)")
}
