rider:
	@go build -o bin/rider rider/main.go
	@./bin/rider

receiver:
	@go build -o bin/receiver data_receiver/main.go
	@./bin/receiver

hotel:
	@go build -o bin/hotel hotel/main.go
	@./bin/hotel

.PHONY: rider, hotel


