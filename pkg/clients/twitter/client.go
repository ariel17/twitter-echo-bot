package twitter

import (
	"fmt"
	"net/http"

	goTwitter "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

const errorMessage = "twitter API response is not HTTP OK"

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
	ID   int64  `json:"id"`
	Text string `json:"text"`
	ScreenName string `json:"screen_name"`
}

// New creates a new client instance based on the environment.
func New() TwitterClient {
	if configs.IsProduction() {
		return newTwitterClient()
	}
	return &MockTwitterClient{}
}

type twitterClient struct {
	c *goTwitter.Client
}

func newTwitterClient() *twitterClient {
	c := configs.GetOAuth1Config()
	client := c.Client(oauth1.NoContext, configs.GetToken())
	return &twitterClient{
		c: goTwitter.NewClient(client),
	}
}

func (tc *twitterClient) Search(query string) ([]Tweet, error) {
	search, response, err := tc.c.Search.Tweets(&goTwitter.SearchTweetParams{
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %+v", errorMessage, response)
	}
	tweets := []Tweet{}
	for _, tweet := range search.Statuses {
		tweets = append(tweets, Tweet{
			ID:   tweet.ID,
			Text: tweet.Text,
			ScreenName: tweet.User.Name,
		})
	}
	return tweets, nil
}

func (tc *twitterClient) Answer(id int64, text string) error {
	_, response, err := tc.c.Statuses.Update(text, &goTwitter.StatusUpdateParams{
		InReplyToStatusID: id,
	})
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %+v", errorMessage, response)
	}
	return nil
}