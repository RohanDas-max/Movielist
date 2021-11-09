package model

type Movielist struct {
	ID     uint      `json:"id" gorm:"null" binding:"-"`
	Name   string    `json:"name" gorm:"not null"`
	Rating []Ratings `gorm:"foreignKey:Rating_id;references:ID"`
}

type Ratings struct {
	Rating_id int    `json:"rating_id" gorm:"primary_key"`
	Rate      int    `json:"rate"`
	Review    string `json:"review"`
}

type Getallmovies struct {
	ID     int    `json:"id"`
	Name   string `json:"Moviename"`
	Rate   int    `json:"rating"`
	Review string `json:"review"`
}
