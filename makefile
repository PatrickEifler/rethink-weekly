watch:
	watchman -- trigger . build '*.*' --  make build notify

unwatch:
	watchman watch-del .

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rewl.linux  .

mac:
	go build -o rewl

build: linux mac

notify:
	/usr/bin/osascript -e "display notification \"build done\""

