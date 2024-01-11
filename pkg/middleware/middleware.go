package middleware

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	keycloak "github.com/hugocortes/go-keycloak"
// )

// // KeycloakMiddleware returns a Gin middleware for Keycloak authentication
// func KeycloakMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Your Keycloak configuration
// 		keycloakConfig := &keycloak.NewServe{
// 			Realm:           "your-realm",
// 			BaseURL:         "https://your-keycloak-domain/auth",
// 			ClientID:        "your-client-id",
// 			ClientSecret:    "your-client-secret",
// 			EnableDebug:     true, // Set to false in production
// 			SkipTokenVerify: false,
// 		}

// 		// Create Keycloak middleware
// 		middleware := keycloak.(keycloakConfig)

// 		// Validate the token
// 		token, err := middleware.ValidateAccessToken(c.Request)
// 		if err != nil {
// 			fmt.Println(err)
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		// You can use 'token' as needed (e.g., get claims, user info, etc.)

// 		// Proceed to the next middleware or handler
// 		c.Next()
// 	}
// }

// func getTokenFromHeader(req *http.Request) (string, error) {
// 	// Get the token from the Authorization header
// 	tokenString := req.Header.Get("Authorization")

// 	// Check if the Authorization header is present
// 	if tokenString == "" {
// 		return "", fmt.Errorf("Authorization header is missing")
// 	}

// 	// Remove the "Bearer " prefix from the token
// 	return tokenString[7:], nil
// }
