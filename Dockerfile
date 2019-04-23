FROM golang:latest

WORKDIR /go/src/url_lookup
COPY src/main .

RUN go get -v
RUN go test -v .
RUN go install -v

# Copy config.properties
COPY src/main/config.properties /go/bin/