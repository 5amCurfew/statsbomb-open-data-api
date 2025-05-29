package models

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
