package lichess

import "net/http"

type LichessClient struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewLichessClient(token string) *LichessClient {
	return &LichessClient{
		token:      token,
		baseURL:    "https://lichess.org/",
		httpClient: &http.Client{},
	}
}
