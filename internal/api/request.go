package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codytheroux96/mlb-scores/internal/app"
	"github.com/joho/godotenv"
)

var (
	basePath = "https://api.balldontlie.io/mlb/v1"
	apiKey   string
)

func init() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error determining executable path:", err)
		os.Exit(1)
	}

	projectRoot := filepath.Dir(execPath)
	envPath := filepath.Join(projectRoot, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("Error loading .env file from %s: %v\n", envPath, err)
		os.Exit(1)
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Printf("Error: API_KEY not found in %s\n", envPath)
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
