package model

import (
	"go_blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int" json:"cid"`
	Desc     string   `gorm:"type:varchar(200);" json:"desc"`
	Content  string   `gorm:"type:longtext;" json:"content"`
	Img      string   `gorm:"type:varchar(100);" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	// data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// todo 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_ART_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCSE
}

// todo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.ERROR_ART_NOT_EXIST
}

// todo 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCSE
}

// 编辑分类
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
