package main

import (
	"kawa/gradingservice/internal/app/grading/dal"
	"kawa/gradingservice/internal/app/grading/delivery/http"
	"kawa/gradingservice/internal/app/grading/handler"
	"kawa/gradingservice/internal/app/grading/repository"
	"kawa/gradingservice/internal/app/grading/usecase"
	"kawa/gradingservice/pkg/server"

	"github.com/gin-gonic/gin"
)

var (
	gradingRepo    repository.GradingRepository
	gradingUseCase usecase.GradingUseCase
)

func init() {
	// Initialize MongoDB connection
	err := dal.Initialize("mongodb://localhost:27017", "gradingdb", "grades")
	if err != nil {
		panic(err)
	}

	// Initialize repositories and use cases
	gradingRepo = repository.NewGradingMongoDBRepository(dal.GetDatabase())
	gradingUseCase = usecase.NewGradingUseCase(gradingRepo)
}

func main() {
	// Setup Gin router
	router := gin.Default()

	// Setup routes and inject dependencies to handlers
	setupRoutes(router)

	// Start the server
	server.RunServer(router, ":8080")
}

func setupRoutes(router *gin.Engine) {
	// Create handlers and inject dependencies
	gradingHandler := handler.NewGradingHandler(gradingUseCase)
	gradingHTTPHandler := http.NewGradingHTTPHandler(gradingHandler)

	// Define routes
	router.GET("/grades/cursus/:cursusID", gradingHTTPHandler.GetGradesByCursusID)
	router.POST("/grades", gradingHTTPHandler.CreateGrade)
	router.GET("/grades/student/:studentID", gradingHTTPHandler.GetGradesByStudentID)
	router.GET("/grades/class/:classID", gradingHTTPHandler.GetGradesByClass)
	router.GET("/grades/:gradeID", gradingHTTPHandler.GetGradeByID)
	router.PUT("/grades/:gradeID", gradingHTTPHandler.UpdateGrade)
	router.DELETE("/grades/:gradeID", gradingHTTPHandler.DeleteGradeByID)
}
