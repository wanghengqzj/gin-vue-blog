package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/response"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//response.OkWithData(map[string]string{
	//	"id": "xxx",
	//}, c)

	response.FailWithCode(response.SettingsError, c)
}
