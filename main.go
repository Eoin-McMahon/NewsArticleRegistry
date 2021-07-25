package main

import (
    "net/http"
    "github.com/Eoin-McMahon/NewsArticleRegistry/handlers"
    "os"
    "context"
    "log"
    "time"
    "os/signal"
)

func main() {
    logger := log.New(os.Stdout, "News_Atricle_Registry:", log.LstdFlags)

    articleHandler := handlers.NewArticle(logger)

    // Serve multiplexer to handle routing 
    serveMux := http.NewServeMux()
    serveMux.Handle("/", articleHandler)

    // Server attributes
    server := http.Server{
        Addr: ":9090",
        Handler: serveMux,
        IdleTimeout: 120 * time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
    }

    // Start Server
    go server.ListenAndServe()

    // Listen for server termination and handle gracefully
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    signal := <-sigChan
    log.Println("Received terminate, performing graceful shutdown...", signal)

    timeout_ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    server.Shutdown(timeout_ctx)
}
