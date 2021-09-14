package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_simple/Models"
)

func GetUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
