FROM golang:1.17.6 AS builder
WORKDIR /usr/src/app
COPY . .
RUN go get
RUN go build wedding.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /usr/src/app/.env ./
COPY --from=builder /usr/src/app/wedding ./
CMD ["./wedding"]  