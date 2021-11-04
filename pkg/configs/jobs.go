package configs

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	searchQueryKey  = "SEARCH_QUERY"
	greetTextKey    = "GREET_TEXT"
	responseTextKey = "RESPONSE_TEXT"
	jobSecondsKey   = "JOB_SECONDS"
)

var (
	searchQuery             string
	greetText, responseText string
	jobSeconds              int
)

// GetGreetText returns the configured tweet introduction to the tweeting user.
// It will be suffixed with the response text.
func GetGreetText() string {
	return greetText
}

// GetResponseText returns the configured tweet text to use in the automated
// response.
func GetResponseText() string {
	return responseText
}

// GetSearchQuery returns the text to use in search of related tweets to answer.
func GetSearchQuery() string {
	return searchQuery
}

// GetJobSeconds returns the amount of seconds to repeat the execution of job
// in search of responses.
func GetJobSeconds() time.Duration {
	return time.Duration(jobSeconds) * time.Second
}

func loadJobKeys() error {
	searchQuery = os.Getenv(searchQueryKey)
	if searchQuery == "" {
		return errors.New("query cannot be empty")
	}

	greetText = os.Getenv(greetTextKey)

	responseText = os.Getenv(responseTextKey)
	if responseText == "" {
		return errors.New("response cannot be empty")
	}

	seconds, err := strconv.Atoi(os.Getenv(jobSecondsKey))
	if err != nil {
		return errors.New("seconds cannot be empty")
	}
	if seconds <= 0 {
		return errors.New("invalid job seconds value")
	}
	jobSeconds = seconds

	return nil
}

func init() {
	if err := loadJobKeys(); err != nil {
		if IsProduction() {
			panic(err)
		} else {
			log.Printf("failed to parse job keys: %+v", err)
		}
	}
}