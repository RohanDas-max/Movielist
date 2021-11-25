package routes

import (
	"movielist/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/api/movies", controller.GetAllMovies)
	r.POST("/api/movie", controller.PostMovie)
	r.POST("/api/review", controller.PostReview)
	r.GET("/api/:id/movie", controller.GetAmovie)
	r.DELETE("/api/:id/delete", controller.DeleteAmovie)
	r.DELETE("/api/purge", controller.Delete)
	return r
}
