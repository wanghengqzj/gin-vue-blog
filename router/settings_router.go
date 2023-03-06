package router

import (
	"gvb_server/api"
)

func (engine RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	engine.GET("/settings", settingsApi.SettingsInfoView)
	engine.PUT("/settings", settingsApi.SettingsUpdateView)
}
