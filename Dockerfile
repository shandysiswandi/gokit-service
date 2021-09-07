FROM golang:1.17-alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o ./app ./cmd/http/main.go
ENTRYPOINT ["/app/app"]