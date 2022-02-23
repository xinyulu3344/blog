package model

import (
    "blog/utils/errmsg"
    "gorm.io/gorm"
)

type Article struct {
    Category Category `gorm:"foreignKey:cid"`
    gorm.Model
    Title   string `gorm:"type:varchar(100);not null" json:"title"`
    Cid     int    `gorm:"type:int;not null" json:"cid"`
    Desc    string `gorm:"type:varchar(200)" json:"desc"`
    Content string `gorm:"type:longtext" json:"content"`
    Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
    err := db.Create(data).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
    var article Article
    err := db.Preload("Category").Where("cid = ?", id).First(&article).Error
    if err != nil {
        return article, errmsg.ERROR_ARTICLE_NOT_EXIST
    }
    return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
    var articles []Article
    err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, errmsg.ERROR
    }
    return articles, errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
    var article Article
    var maps = make(map[string]interface{})
    maps["title"] = data.Title
    maps["desc"] = data.Desc
    maps["content"] = data.Content
    maps["cid"] = data.Cid
    maps["img"] = data.Img
    err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
    var article Article
    err = db.Where("id = ?", id).Delete(&article).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}
