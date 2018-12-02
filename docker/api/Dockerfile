FROM golang:1.11.2-alpine3.8

RUN apk update && \
    apk upgrade && \
    apk --no-cache add tzdata \ 
    git \ 
    build-base

RUN go get github.com/revel/cmd/revel
