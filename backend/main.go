package main

import (
	"GymShark-Tech-Test/pkg/api"
	"log"
)

const (
	serverAddr = "0.0.0.0:8080"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("cannot initialise server... %s", err)
	}

	// main server
	err = server.Start(serverAddr)
	if err != nil {
		log.Fatalf("cannot initialise server... %s", err)
	}
}
