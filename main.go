package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"golang-user-service/config/database"
	userController "golang-user-service/src/account/controllers"
	userRepo "golang-user-service/src/account/repositories"
	userService "golang-user-service/src/account/services"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	errlogfile, _ := os.Create("error.log")
	gin.DefaultErrorWriter = errlogfile

	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	database.Init()
	db := database.GetDB()
	userRepo := userRepo.CreatePersonRepo(db)
	userService := userService.CreatePersonUsecase(userRepo)
	userController.CreateUserHandler(r, userService)
	r.Use(gin.Recovery())
	port := "3302"
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}
