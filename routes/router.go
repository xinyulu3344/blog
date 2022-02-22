package routes

import (
    v1 "blog/api/v1"
    "blog/utils"
    "github.com/gin-gonic/gin"
)

func InitRouter() {
    gin.SetMode(utils.AppMode)
    r := gin.Default()
    
    routerV1 := r.Group("api/v1")
    {
        // 用户模块的路由接口
        routerV1.POST("user/add", v1.AddUser)
        routerV1.GET("users", v1.GetUsers)
        routerV1.PUT("user/:id", v1.EditUser)
        routerV1.DELETE("user/:id", v1.DeleteUser)
        
        // 分类模块的路由接口
        routerV1.POST("category/add", v1.AddCategory)
        routerV1.GET("category", v1.GetCategories)
        routerV1.PUT("category/:id", v1.EditCategory)
        routerV1.DELETE("category/:id", v1.DeleteCategory)
        
        
        // 文章模块的路由接口
    }
    r.Run(utils.HttpPort)
}
