package main

import (
	"fmt"

	"task-manager/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
	"task-manager/controllers"
	"os"
)

func main() {
	
	err := godotenv.Load(".env")
if err != nil {
	fmt.Println("⚠️ .env not found, continuing...")
} else {
	fmt.Println("✅ .env loaded")
}

	config.ConnectDB()

	r := gin.Default()


r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:4200",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept",
		},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// IMPORTANT: handle OPTIONS manually
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	// your routes...
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask) 
	r.PUT("/tasks/:id", controllers.UpdateTask)     
	r.DELETE("/tasks/:id", controllers.DeleteTask) 

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}