package user_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	User string `json:"user"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GetAccountByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "2" {
			logrus.Infof("User not found with id=%v", id)
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusOK, Response{User: "somkiat", Name: "Somkiat Pui", Role: "admin"})
	}
}
