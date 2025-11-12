FROM golang:alpine AS builder

RUN apk add git

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/main /

ENTRYPOINT ["/main"]
