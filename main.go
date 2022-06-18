package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang_app/controller"
)

func main() {

	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	userRepo := controller.New()

	r.POST("users", userRepo.CreateUser)
	r.GET("users", userRepo.GetUsers)
	r.GET("users/:id", userRepo.GetUserById)
	r.PUT("users/:id", userRepo.UpdateUser)
	r.DELETE("users/:id", userRepo.DeleteUser)

	return r
}
