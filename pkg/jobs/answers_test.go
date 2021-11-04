package jobs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/twitter-echo-bot/pkg/clients/twitter"
)

func TestAnswer(t *testing.T) {
	testCases := []struct {
		name         string
		isSuccessful bool
		tweets       []twitter.Tweet
		searchErr    error
		answerErr    error
	}{
		{"ok", true, []twitter.Tweet{{ID: 1, Text: "hello!", ScreenName: "ariel17"}}, nil, nil},
		{"failed by search error", false, nil, errors.New("mocked search error"), nil},
		{"failed by answer error", false, []twitter.Tweet{{ID: 1, Text: "hello!"}}, nil, errors.New("mocked answer error")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client = &twitter.MockTwitterClient{
				Tweets:    tc.tweets,
				SearchErr: tc.searchErr,
				AnswerErr: tc.answerErr,
			}
			err := answer()
			assert.Equal(t, tc.isSuccessful, err == nil)
			if tc.searchErr != nil {
				assert.Equal(t, tc.searchErr, err)
			}
			if tc.answerErr != nil {
				assert.Equal(t, tc.answerErr, err)
			}
		})
	}
}

func TestCreateText(t *testing.T) {
	tweet := twitter.Tweet{
		ScreenName: "ariel17",
	}
	testCases := []struct {
		name     string
		greet    string
		response string
		tweet    twitter.Tweet
		expected string
	}{
		{"response with greeting", "hola", "hellou", tweet, "hola @ariel17 hellou"},
		{"response without greeting", "", "hellou", tweet, "@ariel17 hellou"},
		{"response without greeting", "", "hellou", tweet, "@ariel17 hellou"},
		{"greeting without response", "hola", "", tweet, "hola @ariel17"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			text := createText(tc.tweet, tc.greet, tc.response)
			assert.Equal(t, tc.expected, text)
		})
	}
}