

.PHONY: proto
proto:
	sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=4D0C606C42F6A7BD cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/payment/payment.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o payment-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t payment-service:latest
