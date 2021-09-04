package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neazossa/bookstore-user-api/domain/users"
	"github.com/neazossa/bookstore-user-api/services"
	"github.com/neazossa/bookstore-user-api/util/responses"
)

func CreateUser(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := responses.CreateFailedResponse(string(rune(http.StatusBadRequest)), "invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result := services.CreateUser(u)
	if result.Status != "00" {
		if result.Status == "99" {
			c.JSON(http.StatusInternalServerError, result)
			return
		}
		if result.Status != "99" {
			c.JSON(http.StatusBadRequest, result)
			return
		}
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	fmt.Println("Getting users " + c.Param("id"))
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := responses.CreateFailedResponse(string(rune(http.StatusBadRequest)), "invalid user id")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result := services.GetUser(userId)
	if result.Status != "00" {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func FindUser(c *gin.Context) {

}
