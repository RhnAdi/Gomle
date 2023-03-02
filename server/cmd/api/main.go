package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	docs "github.com/RhnAdi/Gomle/cmd/api/docs"
	"github.com/RhnAdi/Gomle/pkg/app"
	database "github.com/RhnAdi/Gomle/pkg/db"
	PostHandler "github.com/RhnAdi/Gomle/pkg/http/handler/post"
	UserHandler "github.com/RhnAdi/Gomle/pkg/http/handler/user"
	"github.com/RhnAdi/Gomle/pkg/http/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	args := os.Args
	op := "server"

	if len(args) > 1 {
		op = args[0]
	}

	if err := run(op); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start, err: %s", err.Error()))
	}
}

// @title Gomle - Social Media App API
// @version 1.0
// @description Rest API Social Media App using Go Language.
// @BasePath /api/v1
// @contact.email raihanadinugroho9g26@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func run(op string) error {
	r := gin.Default()

	// Docs
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.Static("/public", "./public")

	db, err := database.Init()
	defer database.DBClose(db)

	if err != nil {
		log.Fatalf("errorMessage: failed load database, error: %s", err.Error())
	}

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"authorization", "headers", "content-type"}
	r.Use(cors.New(config))

	// V1
	v1 := r.Group("api/v1")

	// User
	userRepo := database.NewUserDB(db)
	userDetailRepo := database.NewProfileDB(db)
	userService := app.NewUserService(userRepo, userDetailRepo)
	userHandler := UserHandler.NewUserHandler(userService)
	routes.User(v1, userHandler)

	// Post
	postRepo := database.NewPostDB(db)
	postService := app.NewPostService(postRepo)
	postHandler := PostHandler.NewPostHandler(postService)
	routes.Post(v1, postHandler)

	if op == "http_test" {
		r.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello Wolrd !!",
			})
		})

		err := r.Run(":8080")
		return err
	}

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Wolrd !!",
		})
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run(":8080")
	return err
}
