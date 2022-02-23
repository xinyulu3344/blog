package v1

import (
    "blog/middleware"
    "blog/model"
    "blog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Login(c *gin.Context) {
    var data model.User
    _ = c.ShouldBindJSON(&data)
    
    var token string
    var code int
    code = model.CheckLogin(data.UserName, data.Password)
    if code == errmsg.SUCCESS {
        token, _ = middleware.SetToken(data.UserName)
    }
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
        "token": token,
    })
}
