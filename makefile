build:
	go build cmd/main.go -o build/rockit

run:
	go run cmd/main.go

clean:
	rm -r build/
