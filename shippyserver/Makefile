.PHONY:run
run: 
	go run ./

.PHONY:docker
docker: 
	sudo docker build -t shippyserver . && sudo docker run -p 50051:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e DB_HOST=datastore:27017 shippyserver

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50051:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e DB_HOST=datastore:27017 shippyserver

# TODO: compile proto in work project
.PHONY:proto
proto:
	protoc -I./.. --go_out=plugins=micro:. ./../proto/consignment/consignment.proto

.DEFAULT_GOAL := run