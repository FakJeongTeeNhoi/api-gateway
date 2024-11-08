package main

import (
	"fmt"
	"github.com/FakJeongTeeNhoi/api-gateway/controller"
	"github.com/FakJeongTeeNhoi/api-gateway/grpcClient"
	"github.com/FakJeongTeeNhoi/api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server...")

	// TODO: Connect to database using gorm

	server := gin.Default()

	grpcClient.NewSpaceClient()

	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	//corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	//
	//server.Use(cors.New(corsConfig))

	server.Use(middleware.SetAccountInfo())

	// TODO: Add routes here
	controller.InitReverseProxyRoutes(server)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}

	// TODO: Add graceful shutdown
}
