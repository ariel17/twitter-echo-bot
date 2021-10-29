package jobs

import (
	"log"
	"strings"

	"github.com/ariel17/twitter-echo-bot/pkg/clients/twitter"
	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

var client twitter.TwitterClient

func answer() error {
	query := configs.GetSearchQuery()
	log.Printf("Searching tweets with query: %s", query)

	tweets, err := client.Search(query)
	if err != nil {
		return err
	}
	log.Printf("Found this amount of tweets: %d", len(tweets))

	for _, tweet := range tweets {
		if err := client.Answer(tweet.ID, configs.GetResponseText()); err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				log.Printf("WARNING: this tweet was already answered: %+v", tweet)
				continue
			}
			return err
		}
		log.Printf("Answered successfully to tweet %+v", tweet)
	}
	return nil
}

func init() {
	client = twitter.New()
}
