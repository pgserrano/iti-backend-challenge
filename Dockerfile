FROM golang:1.13
MAINTAINER "Paulo Serrano <https://github.com/pgserrano>"
ENV PORT 3000
WORKDIR /go/src/github.com/pgserrano/iti-backend-challenge
ADD . .
WORKDIR /go/src/github.com/pgserrano/iti-backend-challenge/cmd/main
RUN go build -o main
ENTRYPOINT ./main
EXPOSE $PORT