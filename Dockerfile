FROM golang:latest

WORKDIR /go/src/main
COPY src/main/. .

RUN go get -d -v
RUN go install -v
