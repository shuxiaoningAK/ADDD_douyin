package conf

import (
	"ADDD_DOUYIN/model"
)

//执行数据迁移
func migration() {
	//自动迁移模式
	//DB.Set("gorm:table_options", "charset=utf8mb4").
	//	AutoMigrate(&model.User{}).
	//	AutoMigrate(&model.Video{}).
	//	AutoMigrate(&model.Comment{})
	if err := DB.AutoMigrate(
		&model.User{},
		&model.Video{},
		&model.Comment{},
		&model.UserFollow{},
		&model.UserFavorite{},
	); err != nil {
		panic(err)
	}
}
