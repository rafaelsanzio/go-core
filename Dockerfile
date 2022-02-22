FROM golang:1.17.3

WORKDIR /go-core

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./cmd/api/main.go

CMD ["./main"]