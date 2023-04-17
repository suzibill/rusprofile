FROM golang:1.19-alpine

WORKDIR /app
COPY ./ ./

RUN apk add --no-cache make

RUN go mod download
RUN make build

EXPOSE 8080

CMD ["./main"]
