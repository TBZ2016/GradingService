package grading

import (
	"fmt"
	"kawa/gradingservice/internal/app/grading/handler"
	"kawa/gradingservice/internal/app/grading/repository"
	"kawa/gradingservice/internal/app/grading/usecase"
	"kawa/gradingservice/pkg/server"
)

// App represents the grading application.
type App struct {
	Server *server.Server
}

// NewApp creates a new instance of the grading application.
func NewApp() *App {
	// Perform any initialization or setup here.
	// For example, initialize configuration, databases, etc.

	// Create repository and use case instances.
	gradingRepo := repository.NewGradingMongoDBRepository( /* pass necessary parameters */ )
	gradingUseCase := usecase.NewGradingUseCase(gradingRepo)
	gradingHandler := handler.NewGradingHandler(gradingUseCase)

	// Create a new server instance.
	serverConfig := server.Config{
		Port: 8080, // Set your desired port number.
		// Add any other server configuration options here.
	}

	server := server.NewServer(serverConfig)

	// Setup routes and inject dependencies into handlers.
	setupRoutes(server, gradingHandler)

	return &App{
		Server: server,
	}
}

// Run starts the grading application.
func (a *App) Run() {
	// Perform any additional startup logic if needed.
	fmt.Println("Grading service is running.")

	// Start the server.
	if err := a.Server.Start(); err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}

// setupRoutes configures the routes for the grading application.
func setupRoutes(s *server.Server, gradingHandler *handler.GradingHandler) {

	s.Router.GET("/grades/cursus/:cursusID", gradingHandler.GetGradesByCursusID)
	s.Router.POST("/grades", gradingHandler.CreateGrade)
	s.Router.GET("/grades/student/:studentID", gradingHandler.GetGradesByStudentID)
	s.Router.GET("/grades/class/:classID", gradingHandler.GetGradesByClass)
	s.Router.GET("/grades/:gradeID", gradingHandler.GetGradeByID)
	s.Router.PUT("/grades/:gradeID", gradingHandler.UpdateGrade)
	s.Router.DELETE("/grades/:gradeID", gradingHandler.DeleteGradeByID)
}
