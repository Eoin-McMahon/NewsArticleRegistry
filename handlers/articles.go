package handlers

import (
    "log"
    "net/http"
    "github.com/Eoin-McMahon/NewsArticleRegistry/data"
)
type Articles struct {
    log *log.Logger
}

func NewArticle(log *log.Logger) *Articles {
    return &Articles{log}
}

func (article *Articles) ServeHTTP(resp_writer http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodGet {
        article.getArticles(resp_writer, req)
        return
    }

    if req.Method == http.MethodPost {
        article.addArticle(resp_writer, req)
        return
    }

    // Catch all
    resp_writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (article *Articles) getArticles(resp_writer http.ResponseWriter, req *http.Request) {
    article.log.Println("Received GET request for articles")
    articleList := data.GetArticles()

    err := articleList.ToJSON(resp_writer)
    if err != nil {
        http.Error(resp_writer, "Unable to marshal JSON", http.StatusInternalServerError)
    }
}


func (article *Articles) addArticle(resp_writer http.ResponseWriter, req *http.Request) {
    article.log.Println("Received POST request for article")


}
