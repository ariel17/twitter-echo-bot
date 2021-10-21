package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

// NewTwitterClient creates a new client with authentication provided.
func NewTwitterClient() *twitter.Client {
	c := configs.GetOAuth1Config()
	client := c.Client(oauth1.NoContext, configs.GetToken())
	return twitter.NewClient(client)
}