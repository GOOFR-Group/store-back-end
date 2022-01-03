FROM golang:1.16-alpine AS builder

RUN apk update && \
    apk add --no-cache make git ca-certificates && \
    update-ca-certificates && \
    apk add --update npm openjdk8-jre &&  \
    npm install @openapitools/openapi-generator-cli -g

WORKDIR /app

COPY . .

RUN make all
RUN openapi-generator-cli generate -i /app/internal/oapi/store.yaml -g html2 -o /app/tools/swagger && \
    mv /app/tools/swagger/index.html /app/tools/swagger/docs.html

FROM alpine:3.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/dist/store /go/bin/store
COPY --from=builder /app/dist/store-dev /go/bin/store-dev
COPY --from=builder /app/tools/swagger /go/static/swagger

ENTRYPOINT ["/go/bin/store"]