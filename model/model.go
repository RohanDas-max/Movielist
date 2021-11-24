package model

import "gorm.io/gorm"

type Schema interface {
	Movielist()
}

type Movielist struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Rating []Ratings
}

type Ratings struct {
	gorm.Model
	Rate        int    `json:"rate"`
	Review      string `json:"review"`
	MovielistID uint
}

type Getallmovies struct {
	ID     int    `json:"id"`
	Name   string `json:"Moviename"`
	Rate   int
	Review string
}
