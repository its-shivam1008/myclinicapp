package main

import (
    "fmt"
    "net/http"
    "myclinic/internal/routes"
    "myclinic/internal/database"

    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Starting gin server...")
    database.LoadEnv()
    db := database.InitDB()

    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Health Ok"})
    })

    routes.SetupRoutes(r, db)

    fmt.Println("Server running at http://localhost:8080")
    r.Run()
}
