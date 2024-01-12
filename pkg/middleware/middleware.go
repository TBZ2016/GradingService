// middleware/keycloak_service.go

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// KeycloakService handles interactions with Keycloak for authentication and authorization.
type KeycloakService struct {
	IssuerURL    string
	ClientID     string
	ClientSecret string
	Realm        string

	config   *oauth2.Config
	verifier *oidc.IDTokenVerifier
}

// NewKeycloakService initializes a new KeycloakService instance.
func NewKeycloakService(issuerURL, clientID, clientSecret, realm string) *KeycloakService {
	ks := &KeycloakService{
		IssuerURL:    issuerURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Realm:        realm,
	}

	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to create provider: %v", err))
	}

	ks.config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	ks.verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

	return ks
}

// CheckTokenAndRoles checks the validity of the token and the required roles using the Keycloak introspection endpoint.
func (ks *KeycloakService) CheckTokenAndRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ks.extractTokenFromHeader(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Validate token using introspection endpoint
		introspectionResponse, err := ks.introspectToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
			c.Abort()
			return
		}

		if !introspectionResponse.Active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the user has the required roles
		if !ks.checkRoles(c, roles, introspectionResponse.RealmAccess.Roles) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}
	}
}

func (ks *KeycloakService) extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	// Trim any leading or trailing whitespaces
	return strings.TrimSpace(authHeader), nil
}

func (ks *KeycloakService) introspectToken(tokenString string) (*introspectionResponse, error) {
	introspectionURL := ks.IssuerURL + "/protocol/openid-connect/token/introspect"
	client := &http.Client{}

	data := url.Values{}
	data.Set("token", tokenString)
	data.Set("client_id", ks.ClientID)
	data.Set("client_secret", ks.ClientSecret)
	data.Set("token_type_hint", "access_token")

	req, err := http.NewRequest("POST", introspectionURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var introspectionResponse introspectionResponse

	if err := json.NewDecoder(resp.Body).Decode(&introspectionResponse); err != nil {
		return nil, err
	}

	if introspectionResponse.Error != "" {
		return nil, fmt.Errorf("introspection error: %s", introspectionResponse.Error)
	}

	return &introspectionResponse, nil
}

// checkRoles checks if the user has the required roles.
func (ks *KeycloakService) checkRoles(c *gin.Context, requiredRoles, userRoles []string) bool {
	// Check if the user has at least one of the required roles
	for _, requiredRole := range requiredRoles {
		if contains(userRoles, requiredRole) {
			return true
		}
	}

	// Respond with "Insufficient permissions" if roles are not sufficient
	c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
	c.Abort()
	return false
}

// contains checks if a string is present in a slice of strings.
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// introspectionResponse is a struct to hold the response from the token introspection endpoint.
type introspectionResponse struct {
	Active bool   `json:"active"`
	Error  string `json:"error"`

	// Additional fields based on the provided JSON structure
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
}
