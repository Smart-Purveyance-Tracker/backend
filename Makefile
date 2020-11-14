bin/backend:
	GOOS=linux GOARCH=amd64 go build -ldflags "-linkmode external -extldflags -static" -o bin/backend ./cmd/swagger/

.PHONY: generate
generate:
	 swagger generate server -t api/rest-swagger -f swagger-api/swagger.yml --exclude-main

