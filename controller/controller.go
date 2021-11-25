package controller

import (
	"encoding/json"
	"movielist/database"
	"movielist/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	var Getallmovies *[]model.Getmovies

	// SELECT * FROM ratings FULL OUTER JOIN Movielists on ratings.movielist_id = movielists.id;
	if err := database.DB.Raw("SELECT movielists.id,name,rate,review FROM movielists RIGHT JOIN ratings ON movielists.id = ratings.movielist_id;").Scan(&Getallmovies).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "something went wrong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"DATA": Getallmovies,
		})
	}
}

func GetAmovie(c *gin.Context) {
	var GetMovie *[]model.Getmovies
	userid := c.Param("id")
	id, _ := strconv.Atoi(userid)
	if err := database.DB.Raw("SELECT movielists.id,name,rate,review FROM movielists RIGHT JOIN ratings ON movielists.id = ratings.movielist_id Where movielists.id = ?;", id).Scan(&GetMovie).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "something went wrong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"DATA": GetMovie})
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

func DeleteAmovie(c *gin.Context) {
	var movielist *model.Movielist
	var ratings *model.Ratings
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	if err := database.DB.Where("movielist_id = ?", Id).Delete(&ratings).Error; err == nil {
		database.DB.Where("movielist_id = ?", Id).Delete(&movielist)
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "removed successfully",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}

func Delete(c *gin.Context) {

	if err := database.DB.Raw("TRUNCATE TABLE Ratings").Error; err == nil {

		database.DB.Raw("TRUNCATE TABLE Movielists")
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "database deleted successfully",
		})

	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": err,
		})
	}

}
