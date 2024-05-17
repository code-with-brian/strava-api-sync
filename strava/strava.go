package strava

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func RefreshToken(credentials StravaCredentials) (TokenResponse, error) {
	body, err := json.Marshal(credentials)
	if err != nil {
		return TokenResponse{}, err
	}

	resp, err := http.Post(
		"https://www.strava.com/oauth/token",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return TokenResponse{}, err
	}
	defer resp.Body.Close() // Ensure resource cleanup

	if resp.StatusCode != http.StatusOK {
		return TokenResponse{}, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)

	return tokenResponse, err
}

func FetchActivities(accessToken string) ([]byte, error) {
	url := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?access_token=%s&per_page=30", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Ensure resource cleanup
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	return io.ReadAll(resp.Body) // Read response body efficiently
}
