FROM golang:1.16-alpine AS builder

RUN apk update && \
    apk add --no-cache make git ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY . .

RUN go build -v -o dist/store ./cmd

FROM alpine:3.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/dist/store /go/bin/store
COPY --from=builder /app/internal/oapi/store.yaml /go/static/docs/

ENTRYPOINT ["/go/bin/store"]