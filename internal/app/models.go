package app

type Team struct {
	Name         string
	Abbreviation string
	IsHome       bool
}

type Game struct {
	HomeTeam            Team
	AwayTeam            Team
	HomeScore           int
	AwayScore           int
	HomeStartingPitcher string
	AwayStartingPitcher string
	GameTime            string
	GameStatus          string
	Date                string
}

type Scores struct {
	Date  string
	Games []Game
}
