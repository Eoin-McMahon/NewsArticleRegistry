# NewsArticleRegistry

This is a microservice that defines a simple http server that will host news articles. 
It can be interacted with by using http methods; `GET`, `POST` and `PUT`.

This will be used by BriefMe! later, as the microservice to retrieve articles from.

### ⚡️Start the Server

```bash
$ go run news_article_receiver.go
```

### 🎯 Can be easily interacted with using `curl`

```bash
$ curl localhost:9090 -X GET                    // Get Stored Articles
$ curl localhost:9090 -X POST -d '{data}'       // Post a new article
$ curl localhost:9090/<id> -X PUT -d '{data}'   // Update an existing article
```
