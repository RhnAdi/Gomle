package routes

import (
	PostHandler "github.com/RhnAdi/Gomle/pkg/http/handler/post"
	"github.com/RhnAdi/Gomle/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func Post(r *gin.Engine, h *PostHandler.PostHandler) {
	router := r.Group("/post")
	router.Use(middleware.Auth())
	{
		router.GET("/", h.FindAll)
		router.POST("/", h.Create)
		router.GET("/:id", h.Find)
		router.PUT("/:id", h.Update)
		router.DELETE("/:id", h.Delete)
		router.GET("/mypost", h.FindMyPost)
		router.GET("/dashboard", h.FollowingPosts)
	}
}
