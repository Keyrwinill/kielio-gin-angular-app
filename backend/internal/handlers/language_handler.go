package handlers

import "github.com/gin-gonic/gin"

func GetLanguages(c *gin.Context) {

	c.JSON(200, gin.H{
		"languages": []string{"English", "French", "German"},
	})
}

func GetByName(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"name": name,
	})
}
