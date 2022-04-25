FROM golang:latest
WORKDIR $GOPATH/src/github.com/g8ink
ADD . $GOPATH/src/github.com/g8ink
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./g8ink"]