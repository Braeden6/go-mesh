FROM golang:1.23.5-alpine
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go mod tidy && go build -o main

EXPOSE 80
CMD ["./main"]