package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
}

func (base *Controller) CreateCategory(c *gin.Context) {
	var category
	c.BindJSON()
}

func (base *Controller) GetCategory(c *gin.Context) {

}

func (base *Controller) UpdateCategory(c *gin.Context) {

}
func (base *Controller) DeleteCategory(c *gin.Context) {

}
