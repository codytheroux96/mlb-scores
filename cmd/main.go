package main

import (
	"fmt"
	"os"

	"github.com/codytheroux96/mlb-scores/internal/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected you to specify either 'today' or 'yesterday")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "today":
		scores, err := api.GetScores("2023-10-27")
        if err != nil {
            fmt.Println("Error fetching today's scores:", err)
            os.Exit(1)
        }
        fmt.Println("Today's scores:", scores)
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
		os.Exit(1)
	}
}