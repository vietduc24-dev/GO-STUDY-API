package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "tranvietduc")
	uuid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uuid,
		"users":   []string{"user1", "user2", "user3"},
	})
}
