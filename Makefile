bin/backend:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/backend ./cmd/server/
