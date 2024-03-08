FROM golang:1.19.6
RUN go version
COPY . /tmp/test
RUN cd /tmp/test && GO111MODULE=on go build ./
