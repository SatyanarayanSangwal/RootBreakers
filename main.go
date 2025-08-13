package main

import (
	"github.com/SatyanarayanSangwal/RootBreakers/config"
	"github.com/SatyanarayanSangwal/RootBreakers/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Run(":" + config.DB_PORT)
}
