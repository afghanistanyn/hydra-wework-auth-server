package controllers

import (
	"fmt"
	//"context"
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

type ConsentController struct {
}

func (t *ConsentController) Consent(c *gin.Context) {
	var log = globals.Logger()
	var bcontext = bean.NewApplicationContext(beans.Beans)

	log.Info("Hydra consent flow start ...")

	var consentChallenge string
	if c.Request.Method == http.MethodGet {
		consentChallenge = c.Request.URL.Query().Get("consent_challenge")
		//c.Query("login_challenge")
	} else if c.Request.Method == http.MethodPost {
		consentChallenge = c.PostForm("consent_challenge")
	} else {
		c.String(405, "unsupport method")
		return
	}
	if consentChallenge == "" {
		c.String(501, "not a hydra login request")
		return
	}

	var method string
	if c.Request.TLS == nil {
		method = "http"
	} else {
		method = "https"
	}
	errURL := fmt.Sprintf("%s://%s/error?error=", method, c.Request.Host)

	wxClient := bcontext.Get("weWorkClient").(*wework.Client)
	hydraClient := bcontext.Get("hydraClient").(*SDK.OryHydra)

	getConsentRequestOK, err := hydraClient.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(consentChallenge))
	if err != nil {
		log.Errorf("get consent request from hydra error: %v", err)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
		return
	}


	//now try to accept, ignore <getConsentRequestOK.Payload.Skip>
	acceptConsentRequestParams := admin.NewAcceptConsentRequestParams().WithConsentChallenge(consentChallenge)
	uid := getConsentRequestOK.Payload.Subject
	userInfo, err := wxClient.GetUser(uid)
	if err != nil {
		log.Errorf("get information of uid:%s error: %v", uid, err)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
		return
	}
	grantedUserInfo := map[string]interface{}{
		"userId":      userInfo.UserID,
		"name":        userInfo.Name,
		"mobile":      userInfo.Mobile,
		"position":    userInfo.Position,
		"email":       userInfo.Email,
		"gender":      userInfo.Gender,
		"thumbAvatar": userInfo.ThumbAvatar,
	}

	body := models.HandledConsentRequest{
		GrantedAudience: getConsentRequestOK.Payload.RequestedAudience,
		GrantedScope:    setScopes(getConsentRequestOK.Payload.RequestedScope),
		Remember:        true,
		RememberFor:     3600,
		Session: &models.ConsentRequestSessionData{
			AccessToken: grantedUserInfo,
			IDToken:     grantedUserInfo,
		},
	}
	acceptConsentRequestOK, err := hydraClient.Admin.AcceptConsentRequest(acceptConsentRequestParams.WithConsentChallenge(consentChallenge).WithBody(&body))
	if err != nil {
		//consent err, reject it, for retry the login
		//marshal err as reject request body to reject consent
		body := models.RequestDeniedError {
			Code: 11111,
			Description: fmt.Sprintf("%v", err),
			Debug: "last consent err, you can clean the cookie and reentry the auth flow",
		}
		rejectConsentRequestParams := admin.NewRejectConsentRequestParams().WithConsentChallenge(consentChallenge).WithBody(&body)
		rejectConsentRequestOK, err := hydraClient.Admin.RejectConsentRequest(rejectConsentRequestParams)
		if err != nil {
			c.Redirect(http.StatusFound, fmt.Sprintf("%s%s", errURL, err))
			return
		}else {
			c.Redirect(http.StatusFound, rejectConsentRequestOK.Payload.RedirectTo)
			return
		}
	}else {
		c.Redirect(http.StatusFound, acceptConsentRequestOK.Payload.RedirectTo)
		log.Info("Hydra consent flow end ...")
	}
}

func contains(values []string, s string) bool {
	for _, i := range values {
		if i == s {
			return true
		}
	}
	return false
}

func setScopes(scopes []string) []string {
	if contains(scopes, "openid") {
		return scopes
	}

	r := []string{"openid"}
	r = append(r, scopes...)
	return r
}
