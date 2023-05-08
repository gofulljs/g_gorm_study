package model

import (
	"time"
)

/******sql******
CREATE TABLE `userinfos` (
  `class` varchar(10) NOT NULL COMMENT '班级',
  `no` int(11) NOT NULL COMMENT '学号',
  `name` varchar(10) NOT NULL COMMENT '名称',
  `sex` varchar(1) NOT NULL COMMENT '性别',
  `age` int(11) NOT NULL COMMENT '年龄',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`class`,`no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Userinfos [...]
type Userinfos struct {
	Class      string    `gorm:"primaryKey;column:class;type:varchar(10);not null;comment:'班级'" json:"class"`    // 班级
	No         int       `gorm:"primaryKey;column:no;type:int(11);not null;comment:'学号'" json:"no"`              // 学号
	Name       string    `gorm:"column:name;type:varchar(10);not null;comment:'名称'" json:"name"`                 // 名称
	Sex        string    `gorm:"column:sex;type:varchar(1);not null;comment:'性别'" json:"sex"`                    // 性别
	Age        int       `gorm:"column:age;type:int(11);not null;comment:'年龄'" json:"age"`                       // 年龄
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:null;comment:'创建时间'" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:null;comment:'更新时间'" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Userinfos) TableName() string {
	return "userinfos"
}

// UserinfosColumns get sql column name.获取数据库列名
var UserinfosColumns = struct {
	Class      string
	No         string
	Name       string
	Sex        string
	Age        string
	CreateTime string
	UpdateTime string
}{
	Class:      "class",
	No:         "no",
	Name:       "name",
	Sex:        "sex",
	Age:        "age",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
