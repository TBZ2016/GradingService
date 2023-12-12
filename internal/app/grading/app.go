package grading

import (
	"fmt"
	"kawa/gradingservice/internal/app/grading/dal"
	"kawa/gradingservice/internal/app/grading/handler"
	"kawa/gradingservice/internal/app/grading/repository"
	"kawa/gradingservice/internal/app/grading/usecase"
	"kawa/gradingservice/pkg/server"
)

type App struct {
	Server *server.Server
}

func NewApp() *App {

	gradingRepo := repository.NewGradingRepository(dal.GetDatabase())
	gradingUseCase := usecase.NewGradingUseCase(gradingRepo)
	gradingHandler := handler.NewGradingHandler(gradingUseCase)

	serverConfig := server.Config{
		Port: 8080,
	}

	server := server.NewServer(serverConfig)

	setupRoutes(server, gradingHandler)

	return &App{
		Server: server,
	}
}

func (a *App) Run() {
	fmt.Println("Grading service is running.")

	if err := a.Server.Start(); err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}

func setupRoutes(s *server.Server, gradingHandler *handler.GradingHandler) {
	s.Router.GET("/grades/cursus/:cursusID", gradingHandler.GetGradesByCursusID)
	s.Router.POST("/grades", gradingHandler.CreateGrade)
	s.Router.GET("/grades/student/:studentID", gradingHandler.GetGradesByStudentID)
	s.Router.GET("/grades/class/:classID", gradingHandler.GetGradesByClass)
	s.Router.GET("/grades/:gradeID", gradingHandler.GetGradeByID)
	s.Router.PUT("/grades/:gradeID", gradingHandler.UpdateGrade)
	s.Router.DELETE("/grades/:gradeID", gradingHandler.DeleteGradeByID)
}
