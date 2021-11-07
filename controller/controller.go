package controller

import (
	"encoding/json"
	"movielist/database"
	"movielist/model"

	"github.com/gin-gonic/gin"
)

var movies *model.Movielist
var ratings *model.Ratings
var movielist *model.Movies
var rating *model.Rating

func Hello(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode("gola")
}

func GetAllMovies(c *gin.Context) {
	database.DB.Raw("SELECT name FROM movielists;").Scan(&movielist)
	database.DB.Raw("SELECT rate, review FROM ratings WHERE ratings.rating_id != ? ;", "").Scan(&rating)
	json.NewEncoder(c.Writer).Encode(&movielist)
	json.NewEncoder(c.Writer).Encode(&rating)
}

func PostMovie(c *gin.Context) {

	err := c.Bind(&movies)

	if err != nil {
		json.NewEncoder(c.Writer).Encode(err)
	} else {
		database.DB.Create(&movies)
		database.DB.Create(&ratings)
	}
}

func PostReview(c *gin.Context) {
	err := c.Bind(&ratings)
	if err != nil {
		json.NewEncoder(c.Writer).Encode(err)
	} else {
		database.DB.Create(&ratings)
	}
}
