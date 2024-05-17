package main

import (
	"fmt"
	"os"

	"github.com/code-with-brian/frugal-thinker-sync/strava"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
}

func main() {
	credentials := strava.StravaCredentials{
		ClientID:     os.Getenv("STRAVA_CLIENT_ID"),
		ClientSecret: os.Getenv("STRAVA_SECRET"),
		RefreshToken: os.Getenv("STRAVA_REFRESH_TOKEN"),
		GrantType:    "refresh_token",
	}

	newToken, err := strava.RefreshToken(credentials)
	if err != nil {
		fmt.Println("Error refreshing token:", err)
		os.Exit(1)
	}

	activities, err := strava.FetchActivities(newToken.AccessToken)
	if err != nil {
		fmt.Println("Error fetching activities:", err)
		os.Exit(1)
	}

	fmt.Println("Activities JSON:", string(activities))
}
