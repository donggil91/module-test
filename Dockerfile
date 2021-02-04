FROM golang:1.15.6-buster AS builder
ARG VERSION=dev
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o demo \
  -ldflags="-w -X=main.version=${VERSION}" main.go
FROM scratch
COPY --from=builder /build/demo /demo
ENTRYPOINT [ "/demo" ]