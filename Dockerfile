FROM node:18-alpine as webbuilder

WORKDIR /app

COPY . .

RUN cd frontend && npm ci && npm run build

FROM golang:alpine as builder

RUN apk add -U --no-cache ca-certificates

WORKDIR /app

COPY --from=webbuilder /app/frontend/dist/ /app/server/pkg/public/dist/

COPY . .

RUN cd server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags embed -o web cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server/web /usr/bin/

ENV GIN_MODE=release

EXPOSE 8081

ENTRYPOINT ["web"]