.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/main.go

run: build
	./.bin/bot -c config/dev.yaml

build-image:
	docker build -t owl-morning:v0.1 .

start-container:
	docker run --name owl-morning --env-file .env owl-morning:v0.1
