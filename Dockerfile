FROM golang:1.7.1

ADD ./ /go/src
WORKDIR /go/src