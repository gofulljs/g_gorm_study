package tests

import (
	"fmt"
	"g_gorm_study/model"
	"testing"
)

func TestShowDatas(t *testing.T) {
	var list []*model.Userinfos

	err := model.DB.Find(&list).Error
	if err != nil {
		panic(err)
	}

	for _, user := range list {
		fmt.Println(user)
	}
}
