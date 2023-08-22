FROM golang:alpine as builder

RUN apk add -U --no-cache ca-certificates

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/web /usr/bin/

ENV GIN_MODE=release

EXPOSE 8081

ENTRYPOINT ["web"]