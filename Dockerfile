FROM golang:1.17-alpine AS builder

RUN apk update && \
    apk add --no-cache make git ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY . .

RUN go build -o dist/store ./cmd

FROM alpine:3.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/internal/oapi/store.yaml /go/static/docs/
COPY --from=builder /app/dist/store /go/bin/store

ENTRYPOINT ["/go/bin/store"]