package model

import (
	"fmt"
	"go_blog/utils"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数：", err)
	}
	// 禁用表名的复数形式
	db.SingularTable(true)
	//自动迁移模型
	db.AutoMigrate()
	// 设置链接池中最大闲置连接数量
	db.DB().SetMaxIdleConns(10)
	// 设置数据库中最大数连接数
	db.DB().SetMaxOpenConns(100)
	// 设置连接最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.Close()
}
