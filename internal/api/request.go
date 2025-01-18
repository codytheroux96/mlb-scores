package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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

func GetProvidedDatesScores(date, endDate string) ([]string, error) {
	scores, err := GetScores(date, endDate)
	if err != nil {
		return nil, fmt.Errorf("error fetching today's scores: %w", err)
	}

	var results []string
	for _, game := range scores.Data {
		parsedDate, err := time.Parse(time.RFC3339, game.Date)
		if err != nil {
			results = append(results, fmt.Sprintf("Error parsing date for game: %v", err))
			continue
		}

		var winner string
		if game.HomeTeamData.Runs > game.AwayTeamData.Runs {
			winner = game.HomeTeam.DisplayName
		} else if game.AwayTeamData.Runs > game.HomeTeamData.Runs {
			winner = game.AwayTeam.DisplayName
		} else {
			winner = "Game is not over yet"
		}

		formattedDate := parsedDate.Format("2006-01-02")
		awayTeam := game.AwayTeam.DisplayName
		homeTeam := game.HomeTeam.DisplayName
		awayScore := game.AwayTeamData.Runs
		homeScore := game.HomeTeamData.Runs
		status := game.Status

		results = append(results, fmt.Sprintf(
			"Date: %s\nHome Team: %s\nAway Team: %s\nScore: %d - %d\nStatus: %s\nWinner: %s\n",
			formattedDate, homeTeam, awayTeam, awayScore, homeScore, status, winner))
	}

	return results, nil
}
