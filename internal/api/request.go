package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/codytheroux96/mlb-scores/internal/app"
	"github.com/joho/godotenv"
)

var (
	basePath = "https://api.balldontlie.io/mlb/v1"
	apiKey   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	apiKey = os.Getenv("API_KEY")
}

func GetScores(date, endDate string) (*app.Scores, error) {
	req, err := http.NewRequest("GET", basePath+"/games?dates[]="+date+"&dates[]="+endDate, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var scores app.Scores
	err = json.Unmarshal(body, &scores)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &scores, nil
}

func GetTodaysScores() {

}

func GetYesterdaysScores() {

}
