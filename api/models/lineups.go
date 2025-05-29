package models

import "encoding/json"

type LineUp struct {
	TeamID   int    `json:"team_id"`
	TeamName string `json:"team_name"`
	Lineup   []struct {
		PlayerID       int         `json:"player_id"`
		PlayerName     string      `json:"player_name"`
		PlayerNickname interface{} `json:"player_nickname"`
		JerseyNumber   int         `json:"jersey_number"`
		Country        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Cards     []interface{} `json:"cards"`
		Positions []struct {
			PositionID  int         `json:"position_id"`
			Position    string      `json:"position"`
			From        string      `json:"from"`
			To          interface{} `json:"to"`
			FromPeriod  int         `json:"from_period"`
			ToPeriod    interface{} `json:"to_period"`
			StartReason string      `json:"start_reason"`
			EndReason   string      `json:"end_reason"`
		} `json:"positions"`
	} `json:"lineup"`
}

func ParseLineUps(data []byte) (*[]LineUp, error) {
	var lineUps []LineUp
	err := json.Unmarshal(data, &lineUps)
	if err != nil {
		return nil, err
	}
	return &lineUps, nil
}
