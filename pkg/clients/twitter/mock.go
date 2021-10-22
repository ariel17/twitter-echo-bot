package twitter

// MockTwitterClient provides a handful mock object to be used in test units.
type MockTwitterClient struct {
	Tweets []Tweet
	SearchErr error
	AnswerErr error
}

func (mtc *MockTwitterClient) Search(_ string) ([]Tweet, error) {
	return mtc.Tweets, mtc.SearchErr
}

func (mtc *MockTwitterClient) Answer(_ int64, _ string) error {
	return mtc.AnswerErr
}