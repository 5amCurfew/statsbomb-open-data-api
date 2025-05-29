package models

import "encoding/json"

type Match struct {
	MatchID     int    `json:"match_id"`
	MatchDate   string `json:"match_date"`
	KickOff     string `json:"kick_off"`
	Competition struct {
		CompetitionID   int    `json:"competition_id"`
		CountryName     string `json:"country_name"`
		CompetitionName string `json:"competition_name"`
	} `json:"competition"`
	Season struct {
		SeasonID   int    `json:"season_id"`
		SeasonName string `json:"season_name"`
	} `json:"season"`
	HomeTeam struct {
		HomeTeamID     int         `json:"home_team_id"`
		HomeTeamName   string      `json:"home_team_name"`
		HomeTeamGender string      `json:"home_team_gender"`
		HomeTeamGroup  interface{} `json:"home_team_group"`
		Country        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Managers []struct {
			ID       int         `json:"id"`
			Name     string      `json:"name"`
			Nickname interface{} `json:"nickname"`
			Dob      string      `json:"dob"`
			Country  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"managers"`
	} `json:"home_team"`
	AwayTeam struct {
		AwayTeamID     int         `json:"away_team_id"`
		AwayTeamName   string      `json:"away_team_name"`
		AwayTeamGender string      `json:"away_team_gender"`
		AwayTeamGroup  interface{} `json:"away_team_group"`
		Country        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Managers []struct {
			ID       int         `json:"id"`
			Name     string      `json:"name"`
			Nickname interface{} `json:"nickname"`
			Dob      string      `json:"dob"`
			Country  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"managers"`
	} `json:"away_team"`
	HomeScore      int    `json:"home_score"`
	AwayScore      int    `json:"away_score"`
	MatchStatus    string `json:"match_status"`
	MatchStatus360 string `json:"match_status_360"`
	LastUpdated    string `json:"last_updated"`
	LastUpdated360 string `json:"last_updated_360"`
	Metadata       struct {
		DataVersion         string `json:"data_version"`
		ShotFidelityVersion string `json:"shot_fidelity_version"`
		XyFidelityVersion   string `json:"xy_fidelity_version"`
	} `json:"metadata"`
	MatchWeek        int `json:"match_week"`
	CompetitionStage struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"competition_stage"`
	Stadium struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"stadium"`
	Referee struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"referee"`
}

func ParseMatches(data []byte) (*[]Match, error) {
	var matches []Match
	err := json.Unmarshal(data, &matches)
	if err != nil {
		return nil, err
	}
	return &matches, nil
}
