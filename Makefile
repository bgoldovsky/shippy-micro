.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto

.PHONY:build
build:
	docker-compose build

.PHONY:rm
rm:
	docker-compose rm -f

.PHONY:run
run:
	docker-compose up

.DEFAULT_GOAL := run