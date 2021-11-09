package controller

import (
	"encoding/json"
	"movielist/database"
	"movielist/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var movies *model.Movielist
var ratings *model.Ratings
var Getallmovies *[]model.Getallmovies

func GetAllMovies(c *gin.Context) {
	if err := database.DB.Raw("SELECT movielists.id,movielists.name,ratings.rate,ratings.review FROM movielists inner join ratings on ratings.rating_id=movielists.id;").Scan(&Getallmovies).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "something went wrong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"DATA": Getallmovies})
	}
}

func GetAmovie(c *gin.Context) {
	//!TODO
}

func PostMovie(c *gin.Context) {
	err := c.ShouldBind(&movies)
	if err != nil {
		panic(err)
		// json.NewEncoder(c.Writer).Encode(err)
	} else {
		if err := database.DB.Create(&movies).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"alert": "please provide a id",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "Movie stored",
				"movie":   movies,
			})
		}
	}
}

func PostReview(c *gin.Context) {
	err := c.Bind(&ratings)
	if err != nil {
		json.NewEncoder(c.Writer).Encode(err)
	} else {
		if err := database.DB.Create(&ratings).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "please provide a valid Id",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "rating a review recorder",
			})
		}

	}
}
