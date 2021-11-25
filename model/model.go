package model

type Movielist struct {
	ID     int
	Name   string `json:"name" gorm:"not null"`
	Rating []Ratings
}

type Ratings struct {
	ID          int
	Rate        int    `json:"rate"`
	Review      string `json:"review"`
	MovielistID uint
}

type Getmovies struct {
	ID     int    `json:"id"`
	Name   string `json:"Moviename"`
	Rate   int    `json:"rating"`
	Review string `json:"review"`
}
