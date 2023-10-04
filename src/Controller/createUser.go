package controller

import (
	"github.com/gin-gonic/gin"
	rest_err "github.com/wkauan/meu-primeiro-crud-go/src/Configuration/Rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
