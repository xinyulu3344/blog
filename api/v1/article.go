package v1

import (
    "blog/model"
    "blog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
    var data model.Article
    _ = c.ShouldBindJSON(&data)
    code := model.CreateArticle(&data)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": data,
        "message": errmsg.GetErrMsg(code),
    })
}

// todo 查询单个文章信息

// GetArticles 查询文章列表
func GetArticles(c *gin.Context)  {
    pageSize, _ := strconv.Atoi(c.Query("pagesize"))
    pageNum, _ := strconv.Atoi(c.Query("pagenum"))
    if pageSize == 0 {
        pageSize = 20
    }
    if pageNum == 0 {
        pageNum = 1
    }
    data := model.GetArticles(pageSize, pageNum)
    code := errmsg.SUCCESS
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": data,
        "message": errmsg.GetErrMsg(code),
    })
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context)  {
    var data model.Article
    id, _ := strconv.Atoi(c.Param("id"))
    c.ShouldBindJSON(&data)
    code := model.EditArticle(id, &data)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
    })
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context)  {
    id, _ := strconv.Atoi(c.Param("id"))
    code := model.DeleteArticle(id)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
    })
}

