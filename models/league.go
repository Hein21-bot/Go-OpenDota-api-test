package models

type League struct {
	ID     int64  `json:"leagueid"`
	Name   string `json:"name"`
	Tier   string `json:"tier"`
	Region string `json:"region"`
}
