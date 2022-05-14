package conf

import (
	"douyin/model"
)

//执行数据迁移
func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}).
		AutoMigrate(&model.Video{}).
		AutoMigrate(&model.Comment{})
	DB.Model(&model.Video{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Comment{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")

}
