bin/backend:
	GOOS=linux GOARCH=amd64 go build -o bin/backend ./cmd/swagger/

.PHONY: generate
generate:
	swagger generate server -f swagger-api/swagger.yml