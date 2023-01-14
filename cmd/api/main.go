package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	log.Printf("Starting the Flight API on %s port\n", webPort)

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
