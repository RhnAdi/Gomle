package routes

import (
	PostHandler "github.com/RhnAdi/Gomle/server/pkg/http/handler/post"
	"github.com/RhnAdi/Gomle/server/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func Post(r *gin.RouterGroup, h *PostHandler.PostHandler) {
	router := r.Group("/post")
	router.GET("/", h.FindAll)
	router.GET("/:id", h.Find)
	auth := router.Use(middleware.Auth())
	{
		auth.POST("/", h.Create)
		auth.PUT("/:id", h.Update)
		auth.DELETE("/:id", h.Delete)
		auth.GET("/mypost", h.FindMyPost)
		auth.GET("/dashboard", h.FollowingPosts)
	}
}
