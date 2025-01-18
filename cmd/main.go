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
		endDate := date

		scores, err := api.GetProvidedDateScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching today's scores:", err)
			os.Exit(1)
		}

		for _, score := range scores {
			fmt.Println(score)
		}
	case "yesterday":
		date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		endDate := date

		scores, err := api.GetProvidedDateScores(date, endDate)
		if err != nil {
			fmt.Println("Error fetching today's scores:", err)
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
