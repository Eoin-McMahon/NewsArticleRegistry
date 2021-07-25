package handlers

import (
    "log"
    "net/http"
    "context"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/Eoin-McMahon/NewsArticleRegistry/data"
)

type Articles struct {
    log *log.Logger
}

func NewArticleHandler(log *log.Logger) *Articles {
    return &Articles{log}
}

// handle GET requests
func (article *Articles) GET(resp_writer http.ResponseWriter, req *http.Request) {
    article.log.Println("Received GET request, listing all articles...")
    articleList := data.GetArticles()

    err := articleList.ToJSON(resp_writer)
    if err != nil {
        http.Error(resp_writer, "Unable to marshal JSON", http.StatusInternalServerError)
    }
}

// Handle POST requests
func (article *Articles) POST(resp_writer http.ResponseWriter, req *http.Request) {
    article.log.Println("Received POST request, adding article to store...")

    art := req.Context().Value(ArticleKey{}).(data.Article)

    data.AddArticle(&art)
}

// Handle PUT requests
func (article *Articles) PUT(resp_writer http.ResponseWriter, req *http.Request) {
    // parse URI to retrieve ID
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(resp_writer, "Invalid URI, ID variable is malformed", http.StatusBadRequest)
    }

    article.log.Printf("Received PUT request, updating article with ID: %d", id)

    art := req.Context().Value(ArticleKey{}).(data.Article)

    err = data.UpdateArticle(id, &art)
    if err == data.ErrArticleNotFound {
        http.Error(resp_writer, "Article Not Found", http.StatusNotFound)
        return
    }

    if err != nil {
        http.Error(resp_writer, "Article Not Found", http.StatusInternalServerError)
        return
    }
}

type ArticleKey struct{}

// Middleware for validating requests with data
func (article *Articles) MiddleWareValidateArticle(next_handler http.Handler) http.Handler {
    return http.HandlerFunc(func (resp_writer http.ResponseWriter, req *http.Request) {
        art := data.Article{}

        err := art.FromJSON(req.Body)
        if err != nil {
            article.log.Println("ERROR deserializing article")
            http.Error(resp_writer, "Unable to unmarshal article JSON from request", http.StatusBadRequest)
            return
        }

        // add the article to the contxt
        ctx := context.WithValue(req.Context(), ArticleKey{}, art)
        req = req.WithContext(ctx)

        next_handler.ServeHTTP(resp_writer, req)
    })
}
