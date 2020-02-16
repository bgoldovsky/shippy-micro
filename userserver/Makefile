.PHONY:docker
docker: 
	sudo docker build -t userserver . && sudo docker run -p 50053:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e userserver

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50053:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns userserver

.PHONY:proto
proto:
	protoc -I./.. --go_out=plugins=micro:. ./../api/proto/user/user.proto

.DEFAULT_GOAL := run