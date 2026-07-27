package lichess

import "net/http"

type LichessClient struct {
	Token      string
	BaseURL    string
	HttpClient *http.Client
}

func NewLichessClient(token string) *LichessClient {
	return &LichessClient{
		Token:      token,
		BaseURL:    "https://lichess.org/",
		HttpClient: &http.Client{},
	}
}
