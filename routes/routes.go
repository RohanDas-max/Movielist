package routes

import (
	"movielist/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.New()

	r.GET("/api/hello", controller.Hello)
	r.GET("/api/movies", controller.GetAllMovies)
	r.POST("/api/movie", controller.PostMovie)
	r.POST("/api/review", controller.PostReview)
	return r
}
