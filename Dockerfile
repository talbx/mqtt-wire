FROM golang:1.17 as gobuild

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/notification

FROM scratch
WORKDIR /app
COPY --from=gobuild /app/app .
CMD ["./app"]
