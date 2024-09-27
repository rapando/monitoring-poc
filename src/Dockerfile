FROM golang:alpine

WORKDIR  /app
COPY . .
RUN mkdir -p /app/logs
RUN go build -o main -ldflags="-s -w" main.go
CMD ["./main"]