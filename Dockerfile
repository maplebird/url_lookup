FROM golang:latest

WORKDIR /go/src/url_lookup
COPY src/url_lookup .

RUN go get -v
RUN go test -v .
RUN go install -v

# Copy config.properties
COPY src/url_lookup/config.properties /go/bin/

CMD /go/bin/url_lookup