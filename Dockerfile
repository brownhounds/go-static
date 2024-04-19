FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server

FROM scratch

COPY --from=builder /go/bin/server /go/bin/server

ENTRYPOINT ["/go/bin/server"]
