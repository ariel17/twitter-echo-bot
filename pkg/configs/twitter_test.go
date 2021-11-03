package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTwitterKeys(t *testing.T) {
	testCases := []struct{
		name string
		apiKey string
		apiSecret string
		token string
		tokenSecret string
		isSuccessful bool
	}{
		{"ok", "key", "secret", "token", "secret2", true},
		{"api key cannot be empty", "", "secret", "token", "secret2", false},
		{"api secret cannot be empty", "key", "", "token", "secret2", false},
		{"token cannot be empty", "key", "secret", "", "secret2", false},
		{"token secret cannot be empty", "key", "secret", "token", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldKey := os.Getenv(consumerApiKeyKey)
			oldSecret := os.Getenv(consumerApiSecretKey)
			oldToken := os.Getenv(accessTokenKey)
			oldSecret2 := os.Getenv(accessTokenSecretKey)
			defer func() {
				os.Setenv(consumerApiKeyKey, oldKey)
				os.Setenv(consumerApiSecretKey, oldSecret)
				os.Setenv(accessTokenKey, oldToken)
				os.Setenv(accessTokenSecretKey, oldSecret2)
			}()
			os.Setenv(consumerApiKeyKey, tc.apiKey)
			os.Setenv(consumerApiSecretKey, tc.apiSecret)
			os.Setenv(accessTokenKey, tc.token)
			os.Setenv(accessTokenSecretKey, tc.tokenSecret)

			err := loadTwitterKeys()
			assert.Equal(t, tc.isSuccessful, err == nil)

			if tc.isSuccessful {
				assert.Equal(t, tc.apiKey, consumerApiKey)
				assert.Equal(t, tc.apiSecret, consumerApiSecret)
				assert.Equal(t, tc.token, accessToken)
				assert.Equal(t, tc.tokenSecret, accessTokenSecret)
				assert.NotNil(t, GetToken())
				assert.NotNil(t, GetOAuth1Config())
			}
		})
	}
}