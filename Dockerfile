FROM golang:1.7-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o microservice ./main.go

FROM alpine:3.4

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/microservice .

EXPOSE 8080

CMD ["./microservice"]