#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go get -d -v .
RUN go build -o /app -v -o main .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env
CMD ["./main"]
LABEL Name=testebackend Version=0.0.1
EXPOSE 8080
