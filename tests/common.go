package tests

import "g_gorm_study/model"

const (
	dsn = "root:123456@tcp(127.0.0.1:3306)/gormtests?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	err := model.InitDB(dsn)
	if err != nil {
		panic(err)
	}
}
