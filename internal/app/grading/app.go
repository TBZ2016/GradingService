// grading/app.go

package grading

import (
	"fmt"

	"kawa/gradingservice/internal/app/grading/dal"
	"kawa/gradingservice/internal/app/grading/handler"
	"kawa/gradingservice/internal/app/grading/repository"
	"kawa/gradingservice/internal/app/grading/usecase"
	"kawa/gradingservice/pkg/middleware"
	"kawa/gradingservice/pkg/server"
)

// App represents the grading service application.
type App struct {
	Server        *server.Server
	RequiredRoles map[string][]string // Mapping of endpoint paths to required roles
	Keycloak      *middleware.KeycloakService
}

// NewApp initializes a new grading service application.
func NewApp() *App {
	gradingRepo := repository.NewGradingRepository(dal.GetDatabase())
	gradingUseCase := usecase.NewGradingUseCase(gradingRepo)
	gradingHandler := handler.NewGradingHandler(gradingUseCase)

	serverConfig := server.Config{
		Port: 8081,
	}

	server := server.NewServer(serverConfig)

	// Define required roles for each endpoint
	requiredRoles := map[string][]string{
		"/grades/cursus/:cursusId":   {"student", "teacher", "admin"},
		"/grades":                    {"teacher", "admin"},
		"/grades/student/:studentId": {"student", "teacher", "admin"},
		"/grades/class/:classId":     {"teacher", "admin"},
		"/grades/:gradeId":           {"teacher", "admin"},
	}

	keycloak := middleware.NewKeycloakService("http://localhost:8080/realms/GradingRealm", "grading-service", "NUolJ3LNhLkXzW5hBVb26HgKBoiMiqKN", "GradingRealm")

	// Set up routes with authentication middleware
	setupRoutes(server, gradingHandler, keycloak, requiredRoles)

	return &App{
		Server:        server,
		RequiredRoles: requiredRoles,
		Keycloak:      keycloak,
	}
}

// Run starts the grading service.
func (a *App) Run() {
	fmt.Println("Grading service is running.")

	if err := a.Server.Start(); err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}

func setupRoutes(s *server.Server, gradingHandler *handler.GradingHandler, keycloak *middleware.KeycloakService, requiredRoles map[string][]string) {
	// Use a Gin Router Group for routes with required roles
	authenticatedGroup := s.Router.Group("/").Use()

	for endpoint, roles := range requiredRoles {
		// Route specific to the endpoint
		switch endpoint {
		case "/grades/cursus/:cursusId":
			authenticatedGroup.GET(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.GetGradesByCursusID)
		case "/grades":
			authenticatedGroup.POST(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.CreateGrade)
		case "/grades/student/:studentId":
			authenticatedGroup.GET(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.GetGradesByStudentID)
		case "/grades/class/:classId":
			authenticatedGroup.GET(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.GetGradesByClass)
		case "/grades/:gradeId":
			authenticatedGroup.GET(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.GetGradeByID)
			authenticatedGroup.PUT(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.UpdateGrade)
			authenticatedGroup.DELETE(endpoint, keycloak.CheckTokenAndRoles(roles), gradingHandler.DeleteGradeByID)
		}
	}
}
