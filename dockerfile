FROM golang:1.24.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gos cmd/main.go 
EXPOSE 8080
CMD ["./gos"]
