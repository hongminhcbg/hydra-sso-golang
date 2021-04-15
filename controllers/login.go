package controllers

import (
	"github.com/gin-gonic/gin"
	"hydra-sso/hydra_client"
	"log"
	"net/http"
)

type LoginController struct {
	hydra *hydra_client.HydraClient
}

func NewLoginController(hydra *hydra_client.HydraClient) *LoginController {
	return &LoginController{
		hydra: hydra,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {
	loginChallenge := ctx.Query("login_challenge")
	log.Println("login_challenge = ",loginChallenge)

	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"challenge": loginChallenge,
	})
}

func (c *LoginController) AuthUsernamePassword(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}
	log.Printf("forms = %+v", ctx.Request.Form)
	loginChallenge := ctx.Request.Form["challenge"][0]

	acceptLoginResponse, err := c.hydra.AcceptLoginRequest(loginChallenge, hydra_client.AcceptLoginRequest{Subject: "1"})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusFound, acceptLoginResponse.RedirectTo)
}