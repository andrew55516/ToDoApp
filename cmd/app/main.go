package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	tm := NewTodoManager()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authGroup := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	authGroup.GET("/", func(c *gin.Context) {
		todos := tm.GetAll()
		c.JSON(http.StatusOK, todos)
	})

	authGroup.POST("/create", func(c *gin.Context) {
		reqBody := CreateTodoRequest{}

		err := c.Bind(&reqBody)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		todo := tm.Create(reqBody)

		c.JSON(http.StatusOK, todo)
	})

	authGroup.PATCH("/:id/complete", func(c *gin.Context) {
		id := c.Param("id")

		err := tm.Complete(id)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
	})

	authGroup.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := tm.Remove(id)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
	})

	r.Run()

}
