NAME := rethink-weekly
VERSION := 0.1.2

watch:
	watchman -- trigger . build '*.*' --  make build notify

build_docker:
	docker build -t kureikain/$(NAME):$(VERSION) .

push_docker:
	docker build -t kureikain/$(NAME):$(VERSION) .

run_docker:
	docker run --rm -it kureikain/$(NAME):$(VERSION) /bin/sh -l

unwatch:
	watchman watch-del .

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rewl.linux  .

mac:
	go build -o rewl

build: linux mac

notify:
	/usr/bin/osascript -e "display notification \"build done\""

clean:
	rm -rf rewl.* rethink-weekly
