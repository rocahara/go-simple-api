FROM ubuntu

WORKDIR /myApp
COPY ../../go.mod ./
COPY ../../ ./
ENV PATH $PATH:/usr/local/go/bin
RUN apt-get update && apt-get install -y \
    sudo \
    wget && \
    wget https://dl.google.com/go/go1.17.linux-amd64.tar.gz && \
    sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz && \
    rm go1.17.linux-amd64.tar.gz && \
    go build -ldflags '-s -w' ./cmd/simple-api

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

CMD ["./simple-api"]