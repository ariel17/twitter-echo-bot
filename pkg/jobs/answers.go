package jobs

import (
	"fmt"
	"log"
	"strings"

	"github.com/ariel17/twitter-echo-bot/pkg/clients/twitter"
	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

var (
	client twitter.TwitterClient
	cache map[int64]struct{}
)


func answer() error {
	query := configs.GetSearchQuery()
	log.Printf("Searching tweets with query: %s", query)

	tweets, err := client.Search(query)
	if err != nil {
		return err
	}
	log.Printf("Found this amount of tweets: %d", len(tweets))

	for _, tweet := range tweets {
		if _, found := cache[tweet.ID]; found {
			log.Printf("tweet ID %d already answered. Ignoring it.", tweet.ID)
			continue
		}
		cache[tweet.ID] = struct{}{}

		text := createText(tweet, configs.GetGreetText(), configs.GetResponseText())
		log.Printf("will answer with following text: %s", text)
		if err := client.Answer(tweet.ID, text); err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				log.Printf("WARNING: this tweet was already answered: %+v; error: %+v", tweet, err)
				continue
			}
			return err
		}
		log.Printf("Answered successfully to tweet '%+v' with response '%+v'", tweet, text)
	}
	return nil
}

func createText(tweet twitter.Tweet, greet, response string) string {
	text := fmt.Sprintf("%s @%s %s", greet, tweet.ScreenName, response)
	return strings.Trim(text, " ")
}

func init() {
	client = twitter.New()
	cache = map[int64]struct{}{}
}