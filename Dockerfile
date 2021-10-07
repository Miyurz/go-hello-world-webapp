# syntax=docker/dockerfile:1
FROM arm64v8/golang:1.16 AS builder
WORKDIR /go/src/github.com/michaelanckaert/go-hello-world-webapp
COPY app.go ./
COPY go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
EXPOSE 8080
WORKDIR /root/
COPY --from=builder /go/src/github.com/michaelanckaert/go-hello-world-webapp ./
CMD ["./app"]  
