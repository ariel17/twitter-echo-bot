package configs

import (
	"os"

	"github.com/dghubble/oauth1"
)

const (
	consumerApiKeyKey    = "CONSUMER_API_KEY"
	consumerApiSecretKey = "CONSUMER_API_SECRET"
	accessTokenKey       = "ACCESS_TOKEN"
	accessTokenSecretKey = "ACCESS_TOKEN_SECRET"
)

var (
	consumerApiKey, consumerApiSecret string
	accessToken, accessTokenSecret    string
)

// GetOAuth1Config creates a new OAuth v1 configuration object required to
// create a new HTTP client for API access.
func GetOAuth1Config() *oauth1.Config {
	return oauth1.NewConfig(consumerApiKey, consumerApiSecret)
}

// GetToken creates a new OAuth v1 access token object that identifies the user
// accessing the resources.
func GetToken() *oauth1.Token {
	return oauth1.NewToken(accessToken, accessTokenSecret)
}

func loadTwitterKeys() {
	consumerApiKey = os.Getenv(consumerApiKeyKey)
	consumerApiSecret = os.Getenv(consumerApiSecretKey)
	accessToken = os.Getenv(accessTokenKey)
	accessTokenSecret = os.Getenv(accessTokenSecretKey)
}

func init() {
	loadTwitterKeys()
}