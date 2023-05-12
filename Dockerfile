FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main cmd/rusprofile/main.go

FROM alpine:latest

COPY --from=builder /app/main /main
EXPOSE 8080

CMD ["./main"]
