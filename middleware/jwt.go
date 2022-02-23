package middleware

import (
    "blog/utils"
    "blog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "net/http"
    "strings"
    "time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
    UserName string `json:"username"`
    jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string) (string, int) {
    expireTime := time.Now().Add(10 * time.Hour)
    SetClaims := MyClaims{
        username,
        jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer: "blog",
        },
    }
    reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
    token, err := reqClaim.SignedString(JwtKey)
    if err != nil {
        return "", errmsg.ERROR
    }
    return token, errmsg.SUCCESS
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
    setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
        return key, errmsg.SUCCESS
    } else {
        return nil, errmsg.ERROR
    }
    
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenHeader := c.Request.Header.Get("Authorization")
        code := errmsg.SUCCESS
        if tokenHeader == "" {
            code = errmsg.ERROR_TOKEN_NOT_EXIST
            c.JSON(http.StatusOK, gin.H{
                "code": code,
                "message": errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        checkToken := strings.SplitN(tokenHeader, " ", 2)
        if len(checkToken) != 2 && checkToken[0] != "Bearer" {
            code = errmsg.ERROR_TOKEN_TYPE_WRONG
            c.JSON(http.StatusOK, gin.H{
                "code": code,
                "message": errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        key, tCode := CheckToken(checkToken[1])
        if tCode == errmsg.ERROR {
            code = errmsg.ERROR_TOKEN_WRONG
            c.JSON(http.StatusOK, gin.H{
                "code": code,
                "message": errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        if time.Now().Unix() > key.ExpiresAt {
            code = errmsg.ERROR_TOKEN_RUNTIME
            c.Abort()
            c.JSON(http.StatusOK, gin.H{
                "code": code,
                "message": errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        
        c.Set("username", key.UserName)
        c.Next()
    }
}
