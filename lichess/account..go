package lichess

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (c *LichessClient) GetAccount(ctx context.Context) (*Account, error) {
	// Create request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.BaseURL+"api/account",
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")

	// Retrieve response
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Liches returned status: %s", resp.Status)
	}

	var account Account

	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
