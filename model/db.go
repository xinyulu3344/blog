package model

import (
    "blog/utils"
    "database/sql"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "time"
)

var db *gorm.DB
var err error

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName,
        )
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })
    if err != nil {
        fmt.Println("connect db failed! ", err)
        return
    }
    db.AutoMigrate(&User{}, &Article{}, &Category{})
    var sqlDB *sql.DB
    sqlDB, err = db.DB()
    if err != nil {
        fmt.Println(err)
        return
    }
    // SetMaxIdleConns 设置空闲连接池中连接的最大数量
    sqlDB.SetMaxIdleConns(10)
    
    // SetMaxOpenConns 设置打开数据库连接的最大数量。
    sqlDB.SetMaxOpenConns(100)
    
    // SetConnMaxLifetime 设置了连接可复用的最大时间。
    sqlDB.SetConnMaxLifetime(10 * time.Second)
}
