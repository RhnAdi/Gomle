package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RhnAdi/Gomle/pkg/app"
	database "github.com/RhnAdi/Gomle/pkg/db"
	PostHandler "github.com/RhnAdi/Gomle/pkg/http/handler/post"
	UserHandler "github.com/RhnAdi/Gomle/pkg/http/handler/user"
	"github.com/RhnAdi/Gomle/pkg/http/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func run(op string) error {
	r := gin.Default()

	r.Static("/public", "../../public")

	db, err := database.Init()
	defer database.DBClose(db)

	if err != nil {
		log.Fatalf("errorMessage: failed load database, error: %s", err.Error())
	}

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	// User
	userRepo := database.NewUserDB(db)
	userDetailRepo := database.NewUserDetailDB(db)
	userService := app.NewUserService(userRepo, userDetailRepo)
	userHandler := UserHandler.NewUserHandler(userService)
	routes.User(r, userHandler)

	// Post
	postRepo := database.NewPostDB(db)
	postService := app.NewPostService(postRepo)
	postHandler := PostHandler.NewPostHandler(postService)
	routes.Post(r, postHandler)

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

	err = r.Run(":8080")
	return err
}
