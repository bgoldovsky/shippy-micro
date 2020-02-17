.PHONY:docker
docker: 
	sudo docker build -t userclient . && sudo docker run -p 50042:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e userclient

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50042:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns userclient

.PHONY:proto
proto:
	protoc -I./.. --go_out=plugins=micro:. ./../api/proto/user/user.proto

.DEFAULT_GOAL := run