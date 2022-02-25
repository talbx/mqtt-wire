FROM golang:alpine as gobuild

WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/notification

FROM scratch
WORKDIR /app
COPY --from=gobuild /app/app .
COPY --from=gobuild /app/config.yaml .
COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./app"]
