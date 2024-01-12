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
	err := dal.Initialize("mongodb://localhost:27017/", "GradingService")
	if err != nil {
		log.Fatal("Failed to initialize MongoDB:", err)
	}
	defer dal.Close()

	gradingService := grading.NewApp()

	go func() {
		gradingService.Run()
	}()

	waitForShutdown()

	fmt.Println("Grading service gracefully shut down.")
}

func waitForShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
}
