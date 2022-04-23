FROM golang:latest
WORKDIR $GOPATH/src/github.com/g8url
ADD . $GOPATH/src/github.com/g8url
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./g8url"]