package app

type Team struct {
	Name   string
	isHome bool
}

type Game struct {
	HomeTeam    Team
	AwayTeam    Team
	HomeScore   int
	AwayScore   int
	HomePitcher string
	AwayPitcher string
	GameTime    string
	GameStatus  string
}

type Games struct {
	Date  string
	Games []Game
}
