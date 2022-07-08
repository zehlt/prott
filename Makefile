server:
	go run ./cmd/client
	
client:
	go run ./cmd/server


.PHONY: server client