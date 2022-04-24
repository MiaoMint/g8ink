FROM golang:latest
WORKDIR $GOPATH/src/github.com/g8ink
ADD . $GOPATH/src/github.com/g8ink
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./g8ink"]