package main

import (
    "fmt"
    "net/http"
    "myclinic/internal/routes"
    "myclinic/internal/database"

    "github.com/gin-gonic/gin"

    "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "myclinic/docs"
)

// @title My Clinic API
// @version 1.0
// @description API documentation for Receptionist & Doctor portal
// @host localhost:8080
// @BasePath /

func main() {
    fmt.Println("Starting gin server...")
    database.LoadEnv()
    db := database.InitDB()

    r := gin.Default()

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Health Ok"})
    })

    routes.SetupRoutes(r, db)

    fmt.Println("Server running at http://localhost:8080")
    r.Run()
}
