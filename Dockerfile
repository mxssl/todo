FROM golang:1.15.3-alpine3.12 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/mxssl/todo
COPY . .

# Install external dependcies
RUN apk add --no-cache \
  ca-certificates \
  curl \
  git

# Compile binary
RUN CGO_ENABLED=0 \
  GOOS=$(go env GOHOSTOS) \
  GOARCH=$(go env GOHOSTARCH) \
  go build -v -mod=vendor -o todo cmd/todo-server/main.go

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.12.0
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mxssl/todo/todo /todo
RUN chmod +x todo
CMD ["./todo"]
