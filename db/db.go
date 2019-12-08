package db

import (
    "fmt"
    "os"
    "strings"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "exprorstate-api/entity"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    connectionString := getConnectionString()
    db, err = gorm.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }

    autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

func autoMigration() {
    db.AutoMigrate(&entity.User{})
}

func getParamString(param string, defaultValue string) string {
    env := os.Getenv(param)
    if env != "" {
        return env
    }
    return defaultValue
}

func getConnectionString() string {
    host := getParamString("MYSQL_DB_HOST", "db")
    port := getParamString("MYSQL_PORT", "3306")
    user := getParamString("MYSQL_USER", "admin")
    pass := getParamString("MYSQL_PASSWORD", "adminpass1")
    dbname := getParamString("MYSQL_DB", "user")
    protocol := getParamString("MYSQL_PROTOCOL", "tcp")
    dbargs := getParamString("MYSQL_DBARGS", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
        user, pass, protocol, host, port, dbname, dbargs)
}
