package user

import (
	"fmt"
	userhandler "github.com/adeben33/vehicleParkingApi/pkg/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func User(r *gin.Engine, validate *validator.Validate, APiVersion string) *gin.Engine {
	user := userhandler.UserController{Validate: validate}

	userUrl := r.Group(fmt.Sprintf("/api/%v", APiVersion))
	{
		userUrl.GET("/ping", user.Ping)
		userUrl.POST("/signup", user.SignUp)
		userUrl.POST("/login", user.Login)
	}
	return r
}
