FROM golang:1.19.6
RUN go version
RUN go build ./
