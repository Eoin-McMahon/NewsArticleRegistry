package handlers

import (
    "log"
    "net/http"
    "regexp"
    "github.com/Eoin-McMahon/NewsArticleRegistry/data"
    "strconv"
)
type Articles struct {
    log *log.Logger
}

func NewArticle(log *log.Logger) *Articles {
    return &Articles{log}
}

func (article *Articles) ServeHTTP(resp_writer http.ResponseWriter, req *http.Request) {
    // Handle GET Request
    if req.Method == http.MethodGet {
        article.getArticles(resp_writer, req)
        return
    }

    // Handle POST Request
    if req.Method == http.MethodPost {
        article.addArticle(resp_writer, req)
        return
    }

    // Handle PUT Request
    if req.Method == http.MethodPut {
        // expect an ID in the URI
        path := req.URL.Path

        pattern := `/([0-9]+)`
        regex := regexp.MustCompile(pattern)
        group := regex.FindAllStringSubmatch(path, -1)

        // should only have one group in the match
        if len(group) != 1 {
            article.log.Println("Invalid URI - more than one id")
            http.Error(resp_writer, "Invalid URI", http.StatusBadRequest)
            return
        }

        // should only have 2 capture groups, the whole thing and the ID
        if len(group[0]) != 2 {
            article.log.Println("Invalid URI - more than one capture group")
            http.Error(resp_writer, "Invalid URI", http.StatusBadRequest)
            return
        }

        // Cast to int
        idString := group[0][1]
        id, err := strconv.Atoi(idString)

        if err != nil {
            article.log.Println("Invalid URI - could not cast id to int")
            http.Error(resp_writer, "Invalid URI", http.StatusBadRequest)
            return
        }

        article.log.Printf("Successfully retrieved id: %d", id)

        article.updateArticle(id, resp_writer, req)

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

    art := &data.Article{}

    err := art.FromJSON(req.Body)
    if err != nil {
        http.Error(resp_writer, "Unable to unmarshal JSON from request", http.StatusBadRequest)
    }
    article.log.Printf("Article to store: %#v", art)

    data.AddArticle(art)
}

func (article *Articles) updateArticle(id int, resp_writer http.ResponseWriter, req *http.Request) {
    article.log.Println("Received PUT request for article")

    art := &data.Article{}

    err := art.FromJSON(req.Body)
    if err != nil {
        http.Error(resp_writer, "Unable to unmarshal JSON from request", http.StatusBadRequest)
    }

    article.log.Printf("Article to store: %#v", art)

    err = data.UpdateArticle(id, art)
    if err == data.ErrArticleNotFound {
        http.Error(resp_writer, "Article Not Found", http.StatusNotFound)
        return
    }

    if err != nil {
        http.Error(resp_writer, "Article Not Found", http.StatusInternalServerError)
        return
    }
}
