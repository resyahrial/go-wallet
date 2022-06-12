unit-test:
	go mod verify
	go mod download

	# Run unit test
	go test -v ./... -cover -count=1

inject:
	cd cmd/injector; wire

run:
	go run cmd/main.go