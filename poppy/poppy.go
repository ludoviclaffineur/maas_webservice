package poppy

type Poppy struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CarType     string  `json:"carType"`
	ChargeLevel int     `json:"chargeLevel"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	URLPicture  string  `json:"pictureUrl"`
}

type Poppies []Poppy

var poppiesMap = []Poppy{
	Poppy{
		Id:          1,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 80,
		Latitude:    51.2085,
		Longitude:   4.4119,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394613-21933206-333x268-VW--golf.png",
	},
	Poppy{
		Id:          2,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 20,
		Latitude:    51.2069,
		Longitude:   4.4119,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394613-21933206-333x268-VW--golf.png",
	},
	Poppy{
		Id:          3,
		Name:        "A3 CNG",
		CarType:     "gaz",
		ChargeLevel: 60,
		Latitude:    51.2076,
		Longitude:   4.4151,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394599-21933201-333x268-Audi-A3.png",
	},
	Poppy{
		Id:          4,
		Name:        "A3",
		CarType:     "gaz",
		ChargeLevel: 50,
		Latitude:    51.2104,
		Longitude:   4.4150,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394599-21933201-333x268-Audi-A3.png",
	},
	Poppy{
		Id:          5,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 80,
		Latitude:    51.1986,
		Longitude:   4.3944,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394613-21933206-333x268-VW--golf.png",
	},
	Poppy{
		Id:          6,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 20,
		Latitude:    51.2117,
		Longitude:   4.4329,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394613-21933206-333x268-VW--golf.png",
	},
	Poppy{
		Id:          7,
		Name:        "A3 CNG",
		CarType:     "gaz",
		ChargeLevel: 60,
		Latitude:    51.1989,
		Longitude:   4.4008,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394599-21933201-333x268-Audi-A3.png",
	},
	Poppy{
		Id:          8,
		Name:        "A3",
		CarType:     "gaz",
		ChargeLevel: 50,
		Latitude:    51.2139,
		Longitude:   4.3928,
		URLPicture:  "https://v.fastcdn.co/t/fdbc582e/81328161/1506394599-21933201-333x268-Audi-A3.png",
	},
}
