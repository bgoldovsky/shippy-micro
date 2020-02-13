.PHONY:run
run: 
	go run ./

.PHONY:docker
docker: 
	sudo docker build -t shippyclient . && sudo docker run -p 50052:50051 -it -e MICRO_SERVER_ADDRESS=:50051 shippyclient

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50052:50051 -it -e MICRO_SERVER_ADDRESS=:50051 shippyclient

.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

.DEFAULT_GOAL := run