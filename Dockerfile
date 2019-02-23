FROM golang:1.11-alpine
RUN apk add git
WORKDIR /go/src/github.com/isratmir/itchygo

CMD ["/itchy_go"]