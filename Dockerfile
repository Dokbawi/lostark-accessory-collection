FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
ENV PATH="$PATH:$(go env GOPATH)/bin"
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go install github.com/air-verse/air@latest

COPY .air.toml .

CMD ["air"]