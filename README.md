# Strava API Sync

This project is a Go application that interacts with the Strava API to refresh tokens and fetch activities.

## Prerequisites

- Go 1.16 or later
- A Strava account and API credentials
- A `.env` file with your Strava API credentials

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/code-with-brian/frugal-thinker-sync.git
    cd frugal-thinker-sync
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Create a `.env` file in the root directory of the project and add your Strava API credentials:

    ```env
    STRAVA_CLIENT_ID=your_client_id
    STRAVA_SECRET=your_client_secret
    STRAVA_REFRESH_TOKEN=your_refresh_token
    ```

## Usage

Run the application:

```sh
go run main.go
```

The application will:

1. Load the environment variables from the .env file.
2. Use the Strava API credentials to refresh the access token.
3. Fetch activities from Strava using the new access token.
4. Print the activities in JSON format.

Dependencies
github.com/code-with-brian/frugal-thinker-sync/strava
github.com/joho/godotenv
