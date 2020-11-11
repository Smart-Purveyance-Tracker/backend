bin/backend:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/backend ./cmd/server/

.PHONY: generate
generate:
	swagger generate server -f swagger-api/swagger.yml
