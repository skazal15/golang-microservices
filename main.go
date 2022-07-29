package main

import (
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	cfg, _ := database.ConfigUrl()
	fmt.Println(cfg)
	db := database.NewConnDatabase()
	exerciseService := exercise.NewExerciseUsecase(db)
	userUsecase := user.NewUserUsecase(db)
	r.POST(cfg.Urlregister, userUsecase.Register)
	r.POST(cfg.Urllogin, userUsecase.Login)

	r.GET(cfg.UrlGetExerciseById, middleware.WithJWT(userUsecase), exerciseService.GetExerciseByID)
	r.POST(cfg.UrlCreateExercise, middleware.WithJWT(userUsecase), exerciseService.CreateExercise)
	r.POST(cfg.UrlCreateQuestion, middleware.WithJWT(userUsecase), exerciseService.CreateQuestion)
	r.GET(cfg.UrlCalculateScore, middleware.WithJWT(userUsecase), exerciseService.CalculateUserScore)
	r.POST(cfg.UrlCreateAnswer, middleware.WithJWT(userUsecase), exerciseService.CreateAnswer)
	r.Run(":1234")
}
