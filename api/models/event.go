package models

import "encoding/json"

type Event interface{}

func ParseEvents(data []byte) (*[]Event, error) {
	var events []Event
	err := json.Unmarshal(data, &events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}
