package main

import (
	"net/http"

	"github.com/alexfu00/executor"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/execute", func(c *gin.Context) {
		code := c.PostForm("code")

		output, err := executor.ExecuteGoWithFile(code)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Result": "",
				"Error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Result": output,
			"Error":  "",
		})
	})

	router.Run(":8080")
}
