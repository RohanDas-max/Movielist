package model

import "time"

type Movielist struct {
	ID        int       `json:"id" gorm:"unique, not null = true"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"createdat" gorm:"not null"`
	Rating    int8      `json:"rating" gorm:"not null"`
}

type Movies struct {
	Name   string `json:"name"`
	Rating int8   `json:"rating"`
}
