
build:
	go build -o ./openapi-typegen ./cmd/openapi-typegen/

test:
	go test -count=1 -race -v ./...