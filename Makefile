
prepare:
	dep ensure -v

build: prepare
	go build -o ./bin/go-lightning-log -v

tests:
	go test -v -cover

run:
	go run main.go

build-docker: prepare
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/go-lightning-log_linux
	docker build -t go-lightning-log .

