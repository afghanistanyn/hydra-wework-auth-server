package controllers

//for wework callback and hydra callback(when error)

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
)

type CallbackController struct {
}

func (t *CallbackController) Callback(c *gin.Context) {
	var log = globals.Logger()
	var bcontext = bean.NewApplicationContext(beans.Beans)

	var method string
	if c.Request.TLS == nil {
		method = "http"
	} else {
		method = "https"
	}
	errURL := fmt.Sprintf("%s://%s/error?error=", method, c.Request.Host)

	errmsg := c.Query("error")
	errmsgDetail := c.Query("error_description")
	errmsgHint := c.Query("error_hint")
	code := c.Query("code")
	state := c.Query("state")

	//receive errors comes from wework, redirect to /error endpoint
	if errmsg != "" {
		errURL := fmt.Sprintf("%s://%s/error?error=%s&error_description=%s&error_hint=%s", method, c.Request.Host, errmsg, errmsgDetail, errmsgHint)
		c.Redirect(http.StatusFound, errURL)
		return
	}

	//handler wework callback
	wxClient := bcontext.Get("weWorkClient").(*wework.Client)

	//uid, err := wxClient.GetUserInfo(code)
	uid, err := wxClient.GetUserInfo(code, 1)
	if err != nil {
		log.Errorf("Get user info failed. %v", err)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
		return
	}
	log.Infof("User signed in as wework user:%v", uid)

	//acceptlogin
	hydraClient := bcontext.Get("hydraClient").(*SDK.OryHydra)
	acceptloginRequestParams := admin.NewAcceptLoginRequestParams()
	body := models.HandledLoginRequest{
		Subject:  &uid,
		Remember: true,
	}

	acceptLoginRequestOK, err := hydraClient.Admin.AcceptLoginRequest(acceptloginRequestParams.WithLoginChallenge(state).WithBody(&body))
	if err != nil {
		log.Errorf("accept request of loginChallenge(%s) error: %v", state, err)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
		return
	}

	//redirect to hydra to verify
	log.Infof("acceptLoginRequestOk: %s", acceptLoginRequestOK.Payload)
	c.Redirect(http.StatusMovedPermanently, acceptLoginRequestOK.Payload.RedirectTo)
}
