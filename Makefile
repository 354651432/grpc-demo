proto:msg.proto
	@protoc --go_out=. --go-grpc_out=. $<
	@echo done

client:
	@go run client.go

server:
	@go run server.go
