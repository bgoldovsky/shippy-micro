.PHONY:build
build:
	docker-compose build

.PHONY:rm
rm:
	docker-compose rm -f

.PHONY:run
run:
	docker-compose up

.PHONY:shippyserver
shippyserver: 
	go run ./shippyserver

.PHONY:shippyvessel
shippyvessel: 
	go run ./shippyvessel

.PHONY:shippyclient
shippyclient: 
	go run ./shippyclient ./shippyclient/consignment.json

.DEFAULT_GOAL := run