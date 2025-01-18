package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/codytheroux96/mlb-scores/internal/api"
	"github.com/codytheroux96/mlb-scores/internal/app"
	"github.com/codytheroux96/mlb-scores/internal/ui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected you to specify either 'today' or 'yesterday")
		os.Exit(1)
	}

	input := strings.TrimSpace(os.Args[1])
	input = strings.ToLower(input)

	switch {
	case input == "today":
		date := time.Now().Format("2006-01-02")
		endDate := date

		scores, err := api.GetScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching today's scores:", err)
			os.Exit(1)
		}

		gameData := convertToGameData(scores)
		fmt.Println(ui.RenderTable(gameData))
	case input == "yesterday":
		date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		endDate := date

		scores, err := api.GetScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching yesterday's scores:", err)
			os.Exit(1)
		}

		gameData := convertToGameData(scores)
		fmt.Println(ui.RenderTable(gameData))
	case isValidDate(input):
		date := input
		endDate := date

		scores, err := api.GetScores(date, endDate)
		if err != nil {
			fmt.Printf("Error fetching scores for %s: %v\n", date, err)
			os.Exit(1)
		}

		gameData := convertToGameData(scores)
		fmt.Println(ui.RenderTable(gameData))
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
		os.Exit(1)
	}
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func convertToGameData(scores *app.Scores) []ui.GameData {
	var gameData []ui.GameData
	for _, game := range scores.Data {
		var winner string
		if game.HomeTeamData.Runs > game.AwayTeamData.Runs {
			winner = game.HomeTeam.DisplayName
		} else if game.HomeTeamData.Runs < game.AwayTeamData.Runs {
			winner = game.AwayTeam.DisplayName
		} else {
			winner = "Game is not over yet"
		}

		gameData = append(gameData, ui.GameData{
			Date:      game.Date,
			AwayTeam:  game.AwayTeam.DisplayName,
			HomeTeam:  game.HomeTeam.DisplayName,
			HomeScore: game.HomeTeamData.Runs,
			AwayScore: game.AwayTeamData.Runs,
			Status:    game.Status,
			Winner:    winner,
		})
	}
	return gameData
}
