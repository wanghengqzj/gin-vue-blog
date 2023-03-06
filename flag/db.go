package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func Makemigrations() {
	var err error
	err = global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	if err != nil {
		return
	}
	err = global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	if err != nil {
		return
	}
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.BannerModel{},
		&models.MenuBannerModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败！")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
