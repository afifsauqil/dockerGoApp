FROM golang:1.17-alpine

WORKDIR /app

COPY *.go ./
COPY static ./static

EXPOSE 8080

CMD ["go", "run", "main.go"]
