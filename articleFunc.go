package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ArticlesGet(articles *Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := articles.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

type ArticlePostRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func ArticlePost(post *Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := ArticlePostRequest{}
        c.Bind(&requestBody)

        item := Item{
            Title:       requestBody.Title,
            Description: requestBody.Description,
        }
        post.Add(item)

        c.Status(http.StatusNoContent)
    }
}
