.PHONY:run
run: 
	go run ./

.PHONY:docker
docker: 
	sudo docker build -t shippyvessel . && sudo docker run -p 50052:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e DB_HOST=datastore:27017 shippyvessel

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50052:50051 -it -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=:mdns -e DB_HOST=datastore:27017 shippyvessel
	
.DEFAULT_GOAL := run