FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./server/main.go
RUN chmod +x server
EXPOSE 8080
CMD ["./server/main"]