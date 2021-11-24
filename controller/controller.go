package controller

import (
	"encoding/json"
	"movielist/database"
	"movielist/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	var Getallmovies *[]model.Getallmovies
	if err := database.DB.Raw("SELECT * FROM ratings FULL OUTER JOIN Movielists on ratings.movielist_id = movielists.id;").Scan(&Getallmovies).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "something went wrong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"DATA": Getallmovies})
	}
}

func GetAmovie(c *gin.Context) {
	var Getallmovies *model.Getallmovies
	userid := c.Param("id")
	if err := database.DB.Raw("SELECT * FROM ratings FULL OUTER JOIN Movielists on ratings.movielist_id = ?;", userid).Scan(&Getallmovies).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "something went wrong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"DATA": Getallmovies})
	}

}

func PostMovie(c *gin.Context) {
	var movies *model.Movielist

	err := c.Bind(&movies)
	if err != nil {
		panic(err)
		// json.NewEncoder(c.Writer).Encode(err)
	} else {
		if err := database.DB.Create(&movies).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"alert": "OH NO!",
				"err":   err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "Movie stored",
				"movie":   movies,
			})
			// defer c.Request.Body.Close()
		}
	}
}

func PostReview(c *gin.Context) {
	var ratings *model.Ratings

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
