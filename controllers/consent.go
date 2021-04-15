package controllers

import (
	"github.com/gin-gonic/gin"
	"hydra-sso/hydra_client"
	"log"
	"net/http"
)

type ConsentController struct {
	hydra *hydra_client.HydraClient
}

func NewConsentController(hydra *hydra_client.HydraClient) *ConsentController {
	return &ConsentController{hydra: hydra}
}

func (c *ConsentController) GetConsent(ctx *gin.Context)  {
	consentChallenge := ctx.Query("consent_challenge")
	log.Println("[DB][consent challenge] is ", consentChallenge)
	ctx.HTML(http.StatusOK, "consent.html", gin.H{
		"challenge": consentChallenge,
	})
}

func (c *ConsentController) AuthConsent(ctx *gin.Context)  {
	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}
	log.Printf("forms = %+v", ctx.Request.Form)

	scopes := []string{"openid", "offline_access"}
	if values, ok := ctx.Request.Form["email"]; ok && len(values) > 0 {
		scopes = append(scopes, "email")
	}

	if values, ok := ctx.Request.Form["phone"]; ok && len(values) > 0 {
		scopes = append(scopes, "phone")
	}

	consentChallenge := ctx.Request.Form["challenge"][0]
	acceptConsentRequest := hydra_client.AcceptConsentRequest{
		GrantScope: scopes,
		Remember:   false,
		Session: hydra_client.Session{
			AccessToken: hydra_client.AccessToken{
				Subject: "1",
			},
			IdToken:     hydra_client.IdToken{
				Subject: "1",
			},
		},
	}
	resp, err := c.hydra.AcceptConsentRequest(consentChallenge, acceptConsentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}

	log.Println("[DB] redirect to  = ", resp.RedirectTo)
	ctx.Redirect(http.StatusFound, resp.RedirectTo)
}
