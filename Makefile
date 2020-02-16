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
run: 
	go run ./shippyserver

.DEFAULT_GOAL := run