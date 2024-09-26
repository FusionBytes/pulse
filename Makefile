BINARY=pulse

install-deps:
	go install github.com/abice/go-enum@latest

generate:
	go generate ./...

test: 
	go test -v -cover -covermode=atomic ./...

build:
	go build -o bin/${BINARY} ./cmd/server/main.go

unittest:
	go test -short -v ./...

clean:
	if [ -d bin ] ; then rm -rf bin ; fi

format:
	go fmt ./...

analyze:
	go vet ./...


install-dependencies:
	go mod download 

run:
	go run ./cmd/server/main.go
