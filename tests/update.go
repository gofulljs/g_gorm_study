package tests

import (
	"fmt"
	"g_gorm_study/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	info := &model.Userinfos{
		Class: "一一班",
		No:    3,
	}

	err := model.DB.Model(info).Select(model.UserinfosColumns.No).Updates(&model.Userinfos{
		Class: "一一班",
		No:    4,
	}).Error
	assert.NoError(t, err)

	fmt.Println(info)
}
