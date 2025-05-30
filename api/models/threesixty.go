package models

import "encoding/json"

type ThreeSixty struct {
	EventUUID   string    `json:"event_uuid"`
	VisibleArea []float64 `json:"visible_area"`
	FreezeFrame []struct {
		Teammate bool      `json:"teammate"`
		Actor    bool      `json:"actor"`
		Keeper   bool      `json:"keeper"`
		Location []float64 `json:"location"`
	} `json:"freeze_frame"`
}

func ParseThreeSixties(data []byte) (*[]ThreeSixty, error) {
	var threeSixties []ThreeSixty
	err := json.Unmarshal(data, &threeSixties)
	if err != nil {
		return nil, err
	}
	return &threeSixties, nil
}
