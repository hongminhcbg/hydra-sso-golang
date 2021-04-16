package controllers

import (
	"github.com/gin-gonic/gin"
	"hydra-sso/common"
	"hydra-sso/hydra_client"
	"log"
	"net/http"
	"strings"
)

type UserController struct {
	hydra *hydra_client.HydraClient
}

func NewUserController(hydra *hydra_client.HydraClient) *UserController {
	return &UserController{hydra: hydra}
}

func (c *UserController) GetUserInformation(ctx *gin.Context) {
	tokenArgs := ctx.GetHeader("Authorization")
	args := strings.Split(tokenArgs, " ")
	if len(args) != 2 {
		ctx.JSON(http.StatusBadRequest, "Token invalid")
		return
	}

	respIntrospect, err := c.hydra.Introspect(args[1])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Printf("[DB] introspect success %+v\n", respIntrospect)

	body := make(map[string] string)
	body["name"] = common.UserName
	scopes := strings.Split(respIntrospect.Scope, " ")
	for _, scope := range scopes {
		switch scope {
		case "email":
			body["email"] = common.UserEmail
		case "phone":
			body["phone"] = common.UserPhone
		case "address":
			body["address"] = common.UserAddress
		}
	}

	ctx.JSON(http.StatusOK, body)
}