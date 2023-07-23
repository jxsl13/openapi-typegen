
build:
	go build -o ./openapi-typegen ./cmd/openapi-typegen/

clean:
	rm ./openapi-typegen

test:
	go test -count=1 -race -v ./...


fuzz: fuzz-names


fuzz-names:
	go test --fuzztime 300s -v -count=1  -fuzz=FuzzIntersection ./names/