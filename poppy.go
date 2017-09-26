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
	Mode      string    `schema:"mode"`
	FromLat   float64   `schema:"fromLat"`
	FromLon   float64   `schema:"fromLon"`
	ToLat     float64   `schema:"toLat"`
	ToLon     float64   `schema:"toLon"`
	StartTime time.Time `schema:"startTime"`
	EndTime   time.Time `schema:"endTime"`
}
