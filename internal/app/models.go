package app

type Team struct {
	ID               int    `json:"id"`
	Slug             string `json:"slug"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"display_name"`
	ShortDisplayName string `json:"short_display_name"`
	Name             string `json:"name"`
	Location         string `json:"location"`
	League           string `json:"league"`
	Division         string `json:"division"`
}

type TeamData struct {
	Hits         int   `json:"hits"`
	Runs         int   `json:"runs"`
	Errors       int   `json:"errors"`
	InningScores []int `json:"inning_scores"`
}

type ScoringSummary struct {
	Play      string `json:"play"`
	Inning    string `json:"inning"`
	Period    string `json:"period"`
	AwayScore int    `json:"away_score"`
	HomeScore int    `json:"home_score"`
}

type Game struct {
	ID             int              `json:"id"`
	HomeTeam       Team             `json:"home_team"`
	AwayTeam       Team             `json:"away_team"`
	HomeTeamData   TeamData         `json:"home_team_data"`
	AwayTeamData   TeamData         `json:"away_team_data"`
	Season         int              `json:"season"`
	Postseason     bool             `json:"postseason"`
	Date           string           `json:"date"`
	Venue          string           `json:"venue"`
	Attendance     int              `json:"attendance"`
	ConferencePlay bool             `json:"conference_play"`
	Status         string           `json:"status"`
	Period         int              `json:"period"`
	Clock          int              `json:"clock"`
	DisplayClock   string           `json:"display_clock"`
	ScoringSummary []ScoringSummary `json:"scoring_summary"`
}

type Scores struct {
	Data []Game `json:"data"`
	Meta struct {
		PerPage int `json:"per_page"`
	} `json:"meta"`
}
