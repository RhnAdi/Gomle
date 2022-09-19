package routes

import (
	UserHandler "github.com/RhnAdi/Gomle/pkg/http/handler/user"
	"github.com/RhnAdi/Gomle/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, h *UserHandler.UserHandler) {
	// User
	router := r.Group("/users")
	{
		router.POST("/register", h.Register)
		router.POST("/login", h.Login)

		router.GET("/:id/detail", h.FindProfile)

		// User Account
		account := router.Group("/account")
		account.Use(middleware.Auth())
		account.GET("/", h.MyAccount)

		// User Detail ( Profile )
		detail := router.Group("/detail")
		detail.Use(middleware.Auth())
		detail.PUT("/", h.UpdateProfile)

		// User Following
		follow := router.Group("/follow")
		follow.Use(middleware.Auth())
		follow.GET("/:id", h.Follow)
		follow.GET("/followers", h.Followers)
		follow.GET("/followings", h.Followings)

		// User Upload
		upload := router.Group("/upload")
		upload.Use(middleware.Auth())
		upload.POST("/photo_profile", h.UploadPhotoProfile)
		upload.POST("/banner", h.UploadBanner)
	}
}
