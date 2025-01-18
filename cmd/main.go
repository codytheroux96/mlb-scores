package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codytheroux96/mlb-scores/internal/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected you to specify either 'today' or 'yesterday")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "today":
		date := time.Now().Format("2006-01-02")
		endDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

		scores, err := api.GetScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching today's scores:", err)
			os.Exit(1)
		}

		for _, game := range scores.Data {
			parsedDate, err := time.Parse(time.RFC3339, game.Date)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue
			}

			formattedDate := parsedDate.Format("2006-01-02")
			homeTeam := game.HomeTeam.DisplayName
			awayTeam := game.AwayTeam.DisplayName
			homeScore := game.HomeTeamData.Runs
			awayScore := game.AwayTeamData.Runs
			status := game.Status

			fmt.Printf("Date: %s\nHome Team: %s\nAway Team: %s\nScore: %d - %d\nStatus: %s\n\n",
				formattedDate, homeTeam, awayTeam, homeScore, awayScore, status)
		}
	case "yesterday":
		date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		endDate := time.Now().Format("2006-01-02")

		scores, err := api.GetScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching yesterday's scores:", err)
			os.Exit(1)
		}

		for _, game := range scores.Data {
			parsedDate, err := time.Parse(time.RFC3339, game.Date)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue
			}

			formattedDate := parsedDate.Format("2006-01-02")
			homeTeam := game.HomeTeam.DisplayName
			awayTeam := game.AwayTeam.DisplayName
			homeScore := game.HomeTeamData.Runs
			awayScore := game.AwayTeamData.Runs
			status := game.Status

			fmt.Printf("Date: %s\nHome Team: %s\nAway Team: %s\nScore: %d - %d\nStatus: %s\n\n",
				formattedDate, homeTeam, awayTeam, homeScore, awayScore, status)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
		os.Exit(1)
	}
}
