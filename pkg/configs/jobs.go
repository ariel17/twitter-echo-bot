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
	responseTextKey = "RESPONSE_TEXT"
	jobSecondsKey   = "JOB_SECONDS"
	defaultSeconds  = 60
)

var (
	searchQuery, responseText string
	jobSeconds                int
)

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

func loadJobKeys() {
	searchQuery = os.Getenv(searchQueryKey)
	responseText = os.Getenv(responseTextKey)

	seconds, err := strconv.Atoi(os.Getenv(jobSecondsKey))
	if err != nil {
		log.Printf("job seconds value is not numeric. Using default value: %d", defaultSeconds)
		jobSeconds = defaultSeconds
		return
	}

	if seconds <= 0 {
		panic(errors.New("invalid job seconds value"))
	}
	jobSeconds = seconds
}

func init() {
	loadJobKeys()
}
