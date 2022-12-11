package reddit

import (
	"encoding/base64"
	"time"
)

// Credentials required information in order to get the authorization from reddit.
type Credentials struct {
	UserAgent    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// AuthResponse information about the auth response.
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	ExpiringDate time.Time
}

// Auth The variable necessary for reddit authentication API.
var (
	Auth AuthResponse
)

// basicAuth gets the string used for the API basic auth.
func basicAuth(username string, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
}

// bearerTokenAuth gets the string used for the API bearer token authorization type.
func bearerTokenAuth(token string) string {
	return "Bearer " + token
}
