package model

import (
	"fmt"
	"go_blog/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	// fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=local",

	// 	utils.DbUser,
	// 	utils.DbPassWord,
	// 	utils.DbHost,
	// 	utils.DbPort,
	// 	utils.DbName,
	// )
	//root:123456@tcp(192.144.225.220:3306)/ginblog?charset=utf8&parseTime=True&loc=local
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",

		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	// 禁用表名的复数形式
	db.SingularTable(true)
	//自动迁移模型
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// 设置链接池中最大闲置连接数量
	db.DB().SetMaxIdleConns(10)
	// 设置数据库中最大数连接数
	db.DB().SetMaxOpenConns(100)
	// 设置连接最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.Close()
}
