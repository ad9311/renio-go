ARG GO_VERSION=1.23.2
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o ./bin/server ./cmd/server/main.go

FROM alpine:latest

COPY --from=builder /usr/src/app/bin/server /usr/local/bin/
COPY --from=builder /usr/src/app/db/migrations/ /db/migrations/
CMD ["/usr/local/bin/server"]
