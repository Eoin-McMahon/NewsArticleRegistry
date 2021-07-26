FROM golang:latest

ENV GO111MODULE=on

WORKDIR /NewsArticleRegistryService

# External dependencies
COPY go.mod ./
COPY go.sum ./

# Install go dependencies
RUN go mod download

# Service files
COPY *.go ./

EXPOSE 9090

RUN go build -o /NewsArticleRegistry

CMD [ "/NewsArticleRegistry" ]
