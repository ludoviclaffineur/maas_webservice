package main

import (
	"time"
)

type Poppy struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CarType     string  `json:"carType"`
	ChargeLevel int     `json:"chargeLevel"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type Poppies []Poppy

type MaasRequest struct {
	Mode      string    `json:"mode"`
	FromLat   float64   `json:"fromLat"`
	FromLon   float64   `json:"fromLon"`
	ToLat     float64   `json:"toLat"`
	ToLon     float64   `json:"toLon"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
