package main

import (
	"fmt"
	"kawa/gradingservice/internal/app/grading"
	"kawa/gradingservice/internal/app/grading/dal"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize MongoDB connection
	err := dal.Initialize("mongodb://localhost:27017", "gradingdb")
	if err != nil {
		log.Fatal("Failed to initialize MongoDB:", err)
	}

	// Close MongoDB connection on application shutdown
	defer dal.Close()

	// Create grading service instance
	gradingService := grading.NewApp()

	// Start the grading service
	go func() {
		if err := gradingService.Run(); err != nil {
			log.Fatal("Failed to start grading service:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the service
	waitForShutdown()

	fmt.Println("Grading service gracefully shut down.")
}

func waitForShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sigCh
}
