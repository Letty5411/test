FROM golang:1.19.6
RUN go version
RUN GO111MODULE=on go build ./
