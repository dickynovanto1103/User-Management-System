FROM golang:1.13.5-alpine
WORKDIR /go/src/github.com/dickynovanto1103/User-Management-System/
COPY .. .

RUN apk update && apk add git
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only -v
RUN go install ./...

CMD ["tcpserver"]