package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

func LoggerHandlerFunc() gin.HandlerFunc {
    filePath := "log/blog"
    logger := logrus.New()
    src, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        fmt.Println("err: ", err)
        logger.Out = os.Stdout
    }
    logger.Out = src
    logger.SetLevel(logrus.DebugLevel)
    logWriter, _ := rotatelogs.New(filePath + "%Y%m%d.log", rotatelogs.WithMaxAge(7 * 24 * time.Hour), rotatelogs.WithRotationTime(24 * time.Hour))
    writeMap := lfshook.WriterMap{
        logrus.InfoLevel: logWriter,
        logrus.FatalLevel: logWriter,
        logrus.DebugLevel: logWriter,
        logrus.WarnLevel: logWriter,
        logrus.ErrorLevel: logWriter,
        logrus.PanicLevel: logWriter,
    }
    Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
    logger.AddHook(Hook)
    return func(c *gin.Context) {
        startTime := time.Now()
        c.Next()
        spendTime := fmt.Sprintf("%d ms", time.Since(startTime).Milliseconds())
        hostname, err := os.Hostname()
        if err != nil {
            hostname = "unknown"
        }
        statusCode := c.Writer.Status()
        clientIp := c.ClientIP()
        userAgent := c.Request.UserAgent()
        dataSize := c.Writer.Size()
        if dataSize < 0 {
            dataSize = 0
        }
        method := c.Request.Method
        path := c.Request.RequestURI
        
        entry := logger.WithFields(logrus.Fields{
            "clientIp":  clientIp,
            "hostname":  hostname,
            "method":    method,
            "path":      path,
            "status":    statusCode,
            "spendTime": spendTime,
            "dataSize":  dataSize,
            "userAgent": userAgent,
        })
        if len(c.Errors) > 0 {
            entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
        }
        if statusCode >= 500 {
            entry.Error()
        } else if statusCode >= 400 {
            entry.Warn()
        } else {
            entry.Info()
        }
    }
}
