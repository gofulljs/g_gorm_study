package tests

import (
	"fmt"
	"g_gorm_study/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// &{一一班 1 zhang 男 6 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}
// &{一一班 2 zhao 男 6 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}
// &{一一班 3 qian 男 7 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}
// &{一一班 4 sun 女 6 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}
// &{一二班 1 zhang 女 6 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}
// &{一二班 2 li 男 7 2023-05-08 21:43:04 +0800 CST 2023-05-08 21:43:06 +0800 CST}
// &{一二班 3 zhou 男 7 2023-05-08 21:43:04 +0800 CST 2023-05-08 21:43:06 +0800 CST}
// &{一二班 4 wu 男 7 2023-05-08 21:43:04 +0800 CST 2023-05-08 21:43:06 +0800 CST}

// 根据主键获取
func TestQueryByPrimary00(t *testing.T) {
	// &{一一班 3 qian 男 7 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}

	info := &model.Userinfos{
		Class: "一一班",
		No:    3,
	}

	err := model.DB.Take(info).Error
	assert.NoError(t, err)

	fmt.Println(info)
}

func TestQueryByPrimary01(t *testing.T) {
	// &{一一班 3 qian 男 7 2023-05-08 21:40:22 +0800 CST 2023-05-08 21:40:24 +0800 CST}

	info := &model.Userinfos{
		Class: "一一班",
		No:    3,
	}

	err := model.DB.Model(info).Scan(info).Error
	assert.NoError(t, err)

	fmt.Println(info)
}

// 只选取一班的随机一个学生, 防止受主条件影响
func TestQueryByBase01(t *testing.T) {
	info := &model.Userinfos{
		Class: "一一班",
		No:    4,
	}

	err := model.DB.Model(&model.Userinfos{}).Where("class = ?", info.Class).Scan(info).Error
	assert.NoError(t, err)

	fmt.Println(info)
}

// 复用结构体
func TestQueryByBase02(t *testing.T) {
	info := &model.Userinfos{
		Class: "一一班",
		No:    4,
	}

	err := model.DB.Table(info.TableName()).Where("class = ?", info.Class).Scan(info).Error
	assert.NoError(t, err)

	fmt.Println(info)
}
