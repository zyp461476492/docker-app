FROM golang:latest

WORKDIR /docker-app

ADD . /docker-app

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.io

RUN go build -o dockerMain ./web 

EXPOSE 8080

ENTRYPOINT  ["./dockerMain"]
