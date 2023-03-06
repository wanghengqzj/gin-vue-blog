package models

import "time"

// UserCollectModel 定义第三张表，记录用户什么时候什么时间收藏了什么文章 用户收藏文章表
type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
