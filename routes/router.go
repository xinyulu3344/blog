package routes

import (
    v1 "blog/api/v1"
    "blog/middleware"
    "blog/utils"
    "github.com/gin-gonic/gin"
)

func InitRouter() {
    gin.SetMode(utils.AppMode)
    r := gin.New()
    r.Use(middleware.LoggerHandlerFunc() , gin.Recovery() )
    
    v1Auth := r.Group("api/v1")
    v1Auth.Use(middleware.JwtToken())
    {
        // 用户模块的路由接口
        v1Auth.GET("user", v1.GetUsers)
        v1Auth.PUT("user/:id", v1.EditUser)
        v1Auth.DELETE("user/:id", v1.DeleteUser)
        
        // 分类模块的路由接口
        v1Auth.POST("category/add", v1.AddCategory)
        v1Auth.PUT("category/:id", v1.EditCategory)
        v1Auth.DELETE("category/:id", v1.DeleteCategory)
        
        // 文章模块的路由接口
        v1Auth.POST("article/add", v1.AddArticle)
        v1Auth.PUT("article/:id", v1.EditArticle)
        v1Auth.DELETE("article/:id", v1.DeleteArticle)
    }
    v1NoAuth := r.Group("api/v1")
    {
        v1NoAuth.POST("user/add", v1.AddUser)
        v1NoAuth.GET("category", v1.GetCategories)
        v1NoAuth.GET("category/:id/articles", v1.GetCategoryArticles)
        v1NoAuth.GET("article/:id", v1.GetArticleInfo)
        v1NoAuth.GET("article", v1.GetArticles)
        v1NoAuth.POST("login", v1.Login)
    }
    r.Run(utils.HttpPort)
}
