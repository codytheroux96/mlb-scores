package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF7CCB"))
	cellStyle   = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF"))
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
		{Title: headerStyle.Render("Date"), Width: 12},
		{Title: headerStyle.Render("Home Team"), Width: 20},
		{Title: headerStyle.Render("Away Team"), Width: 20},
		{Title: headerStyle.Render("Score"), Width: 10},
		{Title: headerStyle.Render("Status"), Width: 15},
		{Title: headerStyle.Render("Winner"), Width: 20},
	}

	var rows []table.Row
	for _, game := range gameData {
		rows = append(rows, table.Row{
			cellStyle.Render(game.Date),
			cellStyle.Render(game.HomeTeam),
			cellStyle.Render(game.AwayTeam),
			cellStyle.Render(fmt.Sprintf("%d - %d", game.HomeScore, game.AwayScore)),
			cellStyle.Render(game.Status),
			cellStyle.Render(game.Winner),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
	)

	t.SetStyles(table.Styles{
		Header: headerStyle,
		Cell:   cellStyle,
	})
	
	return t.View()
}
