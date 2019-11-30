package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    article := New()
    r := gin.Default()
    r.GET("/article", ArticlesGet(article))
    r.POST("/article", ArticlePost(article))

    r.Run() // listen and serve on 0.0.0.0:8080

}
