.PHONY:run
run: 
	go run ./

.PHONY:docker
docker: 
	sudo docker build -t shippy-app . && sudo docker run -p 50051:50051 -it -MICRO_SERVER_ADDRESS=:5003 shippy-app

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50051:50051 -it -e MICRO_SERVER_ADDRESS=:5003 shippy-app

.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

.DEFAULT_GOAL := run