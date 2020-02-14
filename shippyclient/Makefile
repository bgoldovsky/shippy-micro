.PHONY:run
run: 
	go run ./

.PHONY:docker
docker: 
	sudo docker build -t shippyclient . && sudo docker run -p 50041:50051 -it -e MICRO_REGISTRY=:mdns shippyclient

.PHONY:docker-run
docker-run: 
	sudo docker run -p 50041:50051 -it -e MICRO_REGISTRY=:mdns shippyclient

.DEFAULT_GOAL := run