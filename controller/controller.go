package controller

import (
	"encoding/json"
	"movielist/database"
	"movielist/model"

	"github.com/gin-gonic/gin"
)

var movies *model.Movielist
var movielist []*model.Movies

func Hello(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode("gola")
}

func GetAllMovies(c *gin.Context) {
	database.DB.Raw("SELECT name, rating FROM movielists WHERE id != ? ", nil).Scan(&movielist)
	json.NewEncoder(c.Writer).Encode(&movielist)
}

func PostMovie(c *gin.Context) {

	err := c.Bind(&movies)

	if err != nil {
		json.NewEncoder(c.Writer).Encode(err)
	} else {
		database.DB.Create(&movies).Where("Name = ? and Rating != null", nil)

	}
}
