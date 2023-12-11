package main

import (
	"grading-service/internal/app/grading/handler"
	"grading-service/internal/app/grading/repository"
	"grading-service/internal/app/grading/usecase"
	"grading-service/pkg/server"
)

func main() {
	// Setup dependencies
	gradingRepo := repository.NewGradingRepository()         // Implement this
	gradingUseCase := usecase.NewGradingUseCase(gradingRepo) // Implement this
	gradingHandler := handler.NewGradingHandler(gradingUseCase)

	// Setup server
	s := server.NewServer()
	s.AddHandler("/grades/{cursusId}", "GET", gradingHandler.GetByCursusID)

	// Add other routes as needed

	// Run the server
	s.Run()
}
