FROM golang:1.23-alpine

WORKDIR /go-boilerplate-v2

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o go-boilerplate-v2

CMD ["/go-boilerplate-v2/go-boilerplate-v2"]