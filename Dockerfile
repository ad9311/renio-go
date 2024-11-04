ARG GO_VERSION=1.23.2
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run_app ./cmd/server/main.go

FROM alpine:latest

COPY --from=builder /run_app /usr/local/bin/
COPY --from=builder /usr/src/app/db/ /usr/local/bin/db/
COPY --from=builder /usr/src/app/web/ /usr/local/bin/web/

WORKDIR /usr/local/bin

CMD ["run_app"]
