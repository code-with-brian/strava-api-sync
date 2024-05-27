package strava

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
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

func FetchActivities(accessToken string) ([]Activity, error) {
	url := "https://www.strava.com/api/v3/athlete/activities"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Debug print to check the raw JSON response
	fmt.Println("Raw JSON response:", string(body))

	var activities []Activity
	err = json.Unmarshal(body, &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}
