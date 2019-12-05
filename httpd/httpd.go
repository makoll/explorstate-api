package httpd

import (
    "exprorstate-api/article"
    "exprorstate-api/httpd/handler"

    "github.com/gin-gonic/gin"
)

func Exec() {
    article := article.New()
    r := gin.Default()
    r.GET("/db", handler.DbGet())
    r.GET("/article", handler.ArticlesGet(article))
    r.POST("/article", handler.ArticlePost(article))

    r.Run() // listen and serve on 0.0.0.0:8080

}
