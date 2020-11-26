package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mix-go/bean"
	SDK "github.com/ory/hydra/sdk/go/hydra/client"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/ory/hydra/sdk/go/hydra/models"
	"hydra-wework-auth-server/globals"
	"hydra-wework-auth-server/manifest/beans"
	"hydra-wework-auth-server/wework"
	"net/http"
	"strings"
)

const userAgentKeyword = "wework"
const openIDScope = "openid"

type LoginController struct {
}

func isInWework(r *http.Request) bool {
	return strings.Contains(r.UserAgent(), userAgentKeyword)
}

func (t *LoginController) Login(c *gin.Context) {
	var log = globals.Logger()
	var bcontext = bean.NewApplicationContext(beans.Beans)

	log.Info("Hydra lgin flow start ...")

	var loginChallenge string
	if c.Request.Method == http.MethodGet {
		loginChallenge = c.Request.URL.Query().Get("login_challenge")
	} else if c.Request.Method == http.MethodPost {
		loginChallenge = c.PostForm("login_challenge")
	} else {
		c.String(405, "unsupport method")
		return
	}
	if loginChallenge == "" {
		c.String(500, "not a hydra login request")
		return
	}

	wxClient := bcontext.Get("weWorkClient").(*wework.Client)
	hydraClient := bcontext.Get("hydraClient").(*SDK.OryHydra)

	getLoginRequestParams := admin.NewGetLoginRequestParams().WithLoginChallenge(loginChallenge)
	loginRequestOK, err := hydraClient.Admin.GetLoginRequest(getLoginRequestParams)
	if err != nil {
		log.Errorf("get login request from hydra error: %s", err.Error())
	}

	var method string
	if c.Request.TLS == nil {
		method = "http"
	} else {
		method = "https"
	}
	errURL := fmt.Sprintf("%s://%s/error?error=", method, c.Request.Host)

	if loginRequestOK.Payload.Skip == false {
		//need auth, redirect to login ui
		callbackURL := fmt.Sprintf("%s://%s%s", method, c.Request.Host, "/wework/callback")
		var u string
		if isInWework(c.Request) {
			u = wxClient.GetOAuthURL(callbackURL, loginChallenge)
		} else {
			u = wxClient.GetQRConnectURL(callbackURL, loginChallenge)
		}
		log.Infof("redirect to wework for auth:", u)
		c.Redirect(http.StatusFound, u)

	} else {
		//acceptlogin, now redirect to verify and consent

		hydraClient := bcontext.Get("hydraClient").(*SDK.OryHydra)
		acceptloginRequestParams := admin.NewAcceptLoginRequestParams()
		body := models.HandledLoginRequest{
			Subject:  &loginRequestOK.Payload.Subject,
			Remember: true,
		}
		acceptLoginRequestOK, err := hydraClient.Admin.AcceptLoginRequest(acceptloginRequestParams.WithLoginChallenge(loginChallenge).WithBody(&body))
		if err != nil {
			log.Errorf("accept request of loginChallenge(%s) error: %v", loginChallenge, err)
			c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
			return
		}

		log.Infof("acceptLoginRequestOk: %s", acceptLoginRequestOK.Payload)
		c.Redirect(http.StatusMovedPermanently, acceptLoginRequestOK.Payload.RedirectTo)

	}
	log.Info("Hydra login flow end ...")
}
