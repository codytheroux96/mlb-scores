package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/codytheroux96/mlb-scores/internal/app"
)

var (
	basePath = "https://api.balldontlie.io/mlb/v1"
	apiKey   string
)

func init() {
	apiKey = os.Getenv("MLB_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: MLB_API_KEY not set. Please add it to your environment variables.")
		os.Exit(1)
	}
}

func GetScores(date, endDate string) (*app.Scores, error) {
	apiInfoStartDate := "2002-03-31"

	if date < apiInfoStartDate {
		fmt.Println("The API being used does not provide info before opening day in 2002.")
		os.Exit(1)
	}

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
