package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"users.org/internal/srv"
)

func readEnvInt(name string) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' isn't present", name)
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Environment variable '%s' should be an integer value", name)
	}

	return intVal
}

func main() {
	// read Env variables
	port := readEnvInt("SERVER_PORT")

	// Create and start server
	errChan := make(chan error, 1)
	server := srv.NewServer(port)
	go func() {
		if err := server.Start(); err != nil {
			errChan <- err
		}
	}()

	log.Println("Server started on port:", port)

	// Handle the Interrupt signal
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

	select {
	case err := <-errChan:
		log.Println(err)

	case <-exitChan:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Print("Shutting down...")
		if err := server.Stop(ctx); err != nil {
			log.Println("Failed to stop server")
		}
		log.Println("Done")
	}
}
