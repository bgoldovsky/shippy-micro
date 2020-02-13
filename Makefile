.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

.DEFAULT_GOAL := run