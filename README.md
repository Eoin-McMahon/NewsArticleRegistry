# NewsArticleRegistry

Will be used by BriefMe later as the microservice to retrieve articles from

# Usage
### Start the server 
```bash
$ go run main.go
```

### Can be interacted with using `curl`

```bash
$ curl localhost:9090 -XGET // Get Stored Articles
$ curl localhost:9090 -XPOST -d {your_json_data_here}// Post a new article
```
