package main

import (
	"fmt"

	"task-manager/config"
	"task-manager/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	AllowOrigins:     []string{"http://localhost:4200"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	AllowHeaders:     []string{"Origin", "Content-Type"},
	AllowCredentials: true,
}))

routes.SetupRoutes(r)

r.Run(":8080")
	
}