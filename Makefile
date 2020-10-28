# binary file name:  custom-scheduler
# docker image name: myscheduler or piaoliangkb/myscheduler 
BIN_DIR=bin

init:
	mkdir -p $(BIN_DIR)

local: init
	go build -o=$(BIN_DIR)/custom-scheduler 

build-linux: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=${BIN_DIR}/custom-scheduler

image: build-linux
	docker build --no-cache . -t myscheduler

upload-img: image
	/bin/bash tag-push.sh

clean:
	rm -rf bin/

clean-img:
	docker image rm -f myscheduler
