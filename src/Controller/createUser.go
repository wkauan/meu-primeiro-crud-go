package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	validation "github.com/wkauan/meu-primeiro-crud-go/src/Configuration/Validation"
	request "github.com/wkauan/meu-primeiro-crud-go/src/Controller/Model/Request"
	response "github.com/wkauan/meu-primeiro-crud-go/src/Controller/Model/Response"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	} 
	fmt.Println(userRequest)

	response := response.UserResponse {
		ID: "teste",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	fmt.Println(response)
}
