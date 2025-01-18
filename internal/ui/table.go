package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
)

type GameData struct {
	Date      string
	HomeTeam  string
	AwayTeam  string
	HomeScore int
	AwayScore int
	Status    string
	Winner    string
}

func RenderTable(gameData []GameData) string {
	columns := []table.Column{
		{Title: "Date", Width: 12},
		{Title: "Home Team", Width: 20},
		{Title: "Away Team", Width: 20},
		{Title: "Score", Width: 10},
		{Title: "Status", Width: 15},
		{Title: "Winner", Width: 20},
	}

	var rows []table.Row
	for _, game := range gameData {
		rows = append(rows, table.Row{
			game.Date,
			game.AwayTeam,
			game.HomeTeam,
			fmt.Sprintf("%d - %d", game.HomeScore, game.AwayScore),
			game.Status,
			game.Winner,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
	)

	t.SetStyles(table.DefaultStyles())

	return t.View()
}
