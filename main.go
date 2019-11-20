package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	srv := &http.Server{
		//Handler: MainHandler,
		Addr: ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	http.HandleFunc("/", MainHandler)

	log.Println(srv.ListenAndServe())
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! :)")
}
