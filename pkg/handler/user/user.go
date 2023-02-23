package user

import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	userService "github.com/adeben33/vehicleParkingApi/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
	Validate *validator.Validate
}

func (base *UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Ping"})
}

func (base *UserController) SignUp(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = base.Validate.Struct(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRes, errstring, err := userService.SignUpUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": userRes})
}
