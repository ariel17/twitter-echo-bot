package twitter

// MockTwitterClient provides a handful mock object to be used in test units.
type MockTwitterClient struct {
	Tweets    []Tweet
	SearchErr error
	AnswerErr error
}

// Search is a mocked method; just returns the tweet array and error object.
func (mtc *MockTwitterClient) Search(_ string) ([]Tweet, error) {
	return mtc.Tweets, mtc.SearchErr
}

// Answer is a mocked method; just returns the error object.
func (mtc *MockTwitterClient) Answer(_ int64, _ string) error {
	return mtc.AnswerErr
}