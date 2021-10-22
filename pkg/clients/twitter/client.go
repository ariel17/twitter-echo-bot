package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

// TwitterClient represents a high level client that exposes very specific
// behavior. It simply wraps another complex implementation but gives us the
// ability to inject dependencies.
type TwitterClient interface {
	Search(query string) ([]Tweet, error)
	Answer(id int64, text string) error
}

// Tweet is the simplest representation from a real tweet, in order to be able
// to send some answer.
type Tweet struct {
	ID int64 `json:"id"`
	Text string `json:"text"`
}

// NewTwitterClient creates a new client with authentication provided.
func NewTwitterClient() TwitterClient {
	c := configs.GetOAuth1Config()
	client := c.Client(oauth1.NoContext, configs.GetToken())
	return twitter.NewClient(client)
}