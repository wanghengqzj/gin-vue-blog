package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	response.OkWithData(global.Config.SiteInfo, c)
}
