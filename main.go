package main

import (
    "exprorstate-api/httpd"

    "database/sql"
    "fmt"
    "os"
    "strings"

        _ "github.com/go-sql-driver/mysql"
)

func main() {
    connectionString := getConnectionString()
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        fmt.Println(16, err)
    }
    defer db.Close()

    var last_name string
    err = db.QueryRow("SELECT last_name FROM user").Scan(&last_name)
    if err != nil {
        fmt.Println(23, err)
    }
    fmt.Println(25, last_name)
    httpd.Exec()
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
