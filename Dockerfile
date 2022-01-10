# Sample from @citizen428 https://dev.to/citizen428/comment/6cmh
FROM golang:alpine as build
RUN apk add --no-cache ca-certificates
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags '-extldflags "-static"' -o app

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt
COPY --from=build /build/app /app
ENTRYPOINT ["/app"]
