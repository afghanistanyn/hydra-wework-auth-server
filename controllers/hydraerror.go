package controllers

import (
	//"context"
	"github.com/gin-gonic/gin"
	"hydra-wework-auth-server/globals"
	"net/http"
)

type HydraErrorController struct {
}

//may be wework error comes here
func (t *HydraErrorController) HydraError(c *gin.Context) {
	var log = globals.Logger()

	errMsg := c.Request.URL.Query().Get("error")
	errMsgDetail := c.Request.URL.Query().Get("error_description")
	errHint := c.Query("err_hint")
	log.Errorf("error happend:", errMsg, errMsgDetail, errHint)
	errdata := map[string]string{
		"err":       errMsg,
		"errDetail": errMsgDetail,
		"errHint":   errHint,
	}
	c.JSON(http.StatusFound, errdata)
}
