package poppy

type Poppy struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CarType     string  `json:"carType"`
	ChargeLevel int     `json:"chargeLevel"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type Poppies []Poppy
