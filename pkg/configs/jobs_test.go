package configs

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadJobKeys(t *testing.T) {
	testCases := []struct{
		name string
		query string
		greet string
		response string
		seconds string
		isSuccessful bool
	}{
		{"ok", "query", "greet", "response", "99", true},
		{"query cannot be empty", "", "greet", "response", "99", false},
		{"greet can be empty", "query", "", "response", "99", true},
		{"response cannot be empty", "query", "greet", "", "99", false},
		{"seconds cannot be empty", "query", "greet", "response", "", false},
		{"seconds cannot be negative", "query", "greet", "response", "-10", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldQuery := os.Getenv(searchQueryKey)
			oldGreet := os.Getenv(greetTextKey)
			oldResponse := os.Getenv(responseTextKey)
			oldSeconds := os.Getenv(jobSecondsKey)
			defer func() {
				os.Setenv(searchQueryKey, oldQuery)
				os.Setenv(greetTextKey, oldGreet)
				os.Setenv(responseTextKey, oldResponse)
				os.Setenv(jobSecondsKey, oldSeconds)
			}()
			os.Setenv(searchQueryKey, tc.query)
			os.Setenv(greetTextKey, tc.greet)
			os.Setenv(responseTextKey, tc.response)
			os.Setenv(jobSecondsKey, tc.seconds)

			err := loadJobKeys()
			assert.Equal(t, tc.isSuccessful, err == nil)

			if tc.isSuccessful {
				assert.Equal(t, tc.query, GetSearchQuery())
				assert.Equal(t, tc.greet, GetGreetText())
				assert.Equal(t, tc.response, GetResponseText())
				v, _ := strconv.Atoi(tc.seconds)
				assert.Equal(t, time.Duration(v) * time.Second, GetJobSeconds())
			}
		})
	}
}