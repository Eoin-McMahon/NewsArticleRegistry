package main

import (
    "net/http"
    "github.com/Eoin-McMahon/NewsArticleRegistry/handlers"
    "github.com/gorilla/mux"
    "os"
    "context"
    "log"
    "time"
    "os/signal"
)

var address string = "127.0.0.1:9090"

func main() {
    logger := log.New(os.Stdout, "News_Atricle_Registry:", log.LstdFlags)

    articleHandler := handlers.NewArticleHandler(logger)

    // Serve multiplexer to handle routing 
    serveMux := mux.NewRouter()

    // Handle HTTP GET requests
    getRouter := serveMux.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/", articleHandler.GET)

    // Handle HTTP PUT requests, expect an ID in URI
    putRouter := serveMux.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", articleHandler.PUT)
    putRouter.Use(articleHandler.MiddleWareValidateArticle)

    // Handle HTTP POST requests
    postRouter := serveMux.Methods(http.MethodPost).Subrouter()
    postRouter.HandleFunc("/", articleHandler.POST)
    postRouter.Use(articleHandler.MiddleWareValidateArticle)

    // Server attributes
    server := http.Server {
        Addr: address,
        Handler: serveMux,
        IdleTimeout: 120 * time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
    }

    // Start Server in a go routine
    go server.ListenAndServe()
    log.Printf("Server created, listening on %s", address)

    // Listen for manual server termination and handle gracefully
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    signal := <-sigChan
    log.Println("Received terminate, performing graceful shutdown...", signal)

    // Shutdown
    timeout_ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    server.Shutdown(timeout_ctx)
}
