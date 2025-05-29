package models

import "encoding/json"

type Competition struct {
	CompetitionID            int    `json:"competition_id"`
	SeasonID                 int    `json:"season_id"`
	CountryName              string `json:"country_name"`
	CompetitionName          string `json:"competition_name"`
	CompetitionGender        string `json:"competition_gender"`
	CompetitionYouth         bool   `json:"competition_youth"`
	CompetitionInternational bool   `json:"competition_international"`
	SeasonName               string `json:"season_name"`
	MatchUpdated             string `json:"match_updated"`
	MatchUpdated360          string `json:"match_updated_360"`
	MatchAvailable360        string `json:"match_available_360"`
	MatchAvailable           string `json:"match_available"`
}

func ParseCompetitions(data []byte) (*[]Competition, error) {
	var comps []Competition
	err := json.Unmarshal(data, &comps)
	if err != nil {
		return nil, err
	}
	return &comps, nil
}

func ParseCompetition(data []byte) (*Competition, error) {
	var comp Competition
	err := json.Unmarshal(data, &comp)
	if err != nil {
		return nil, err
	}
	return &comp, nil
}
