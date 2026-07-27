package lichess

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ComputerGameOptions struct {
	Level          int
	ClockLimit     int
	ClockIncrement int
	Color          string
}

type ComputerGame struct {
	ID string `json:"id"`
}

func (c *LichessClient) CreateComputerGame(
	ctx context.Context,
	options ComputerGameOptions,
) (*ComputerGame, error) {
	form := url.Values{}
	form.Set("level", strconv.Itoa(options.Level))
	form.Set("clock.limit", strconv.Itoa(options.ClockLimit))
	form.Set("clock.increment", strconv.Itoa((options.ClockIncrement)))
	form.Set("color", options.Color)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.BaseURL+"api/challenge/ai",
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(
			"Lichess returned status: %s",
			resp.Status,
		)
	}
	var game ComputerGame

	err = json.NewDecoder(resp.Body).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
