package main

import (
	"GymShark-Tech-Test/pkg/api"
	"log"
	"os"
)

// const (
// 	serverAddr = "0.0.0.0:8080"
// )

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("cannot initialise server... %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// main server
	err = server.Start(":" + port)
	if err != nil {
		log.Fatalf("cannot initialise server... %s", err)
	}
}
