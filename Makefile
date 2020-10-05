BIN_DIR=bin

init:
	mkdir -p $(BIN_DIR)

local: init
	go build -o=$(BIN_DIR)/myscheduler 

build-linux: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=${BIN_DIR}/myscheduler

image: build-linux
	docker build --no-cache . -t myscheduler

clean:
	rm -rf bin/

clean-img:
	docker image rm -f myscheduler
