package main
 
import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting gin server...")
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Health Ok",
		})
	});
	r.Run();
	fmt.Println("http://localhost:8080")
}