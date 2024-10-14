ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

# Install goose in the builder stage
RUN apk add --no-cache git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Build the Go app
RUN go build -v -o /run-app ./cmd/server/*.go

# Migration stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy the built app and goose binary
COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /go/bin/goose /usr/local/bin/

# Copy migration files
COPY ./migrations /migrations

# Run migrations
CMD ["sh", "-c", "goose -dir /db/migrations postgres \"$DATABASE_URL\" up && run-app"]
