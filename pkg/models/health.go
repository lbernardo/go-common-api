package models

import "time"

type Health struct {
	Success bool   `json:"success"`
	Time    string `json:"time"`
}

func NewHealth() *Health {
	now := time.Now()
	return &Health{
		Time:    now.String(),
		Success: true,
	}
}
