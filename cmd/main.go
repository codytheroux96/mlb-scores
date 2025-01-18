package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/codytheroux96/mlb-scores/internal/api"
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

		scores, err := api.GetProvidedDateScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching today's scores:", err)
			os.Exit(1)
		}

		for _, score := range scores {
			fmt.Println(score)
		}
	case input == "yesterday":
		date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		endDate := date

		scores, err := api.GetProvidedDateScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching yesterday's scores:", err)
			os.Exit(1)
		}

		for _, score := range scores {
			fmt.Println(score)
		}
	case isValidDate(input):
		date := input
		endDate := date

		scores, err := api.GetProvidedDateScores(date, endDate)
		if err != nil {
			fmt.Printf("Error fetching scores for %s: %v\n", date, err)
			os.Exit(1)
		}

		for _, score := range scores {
			fmt.Println(score)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
		os.Exit(1)
	}
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
