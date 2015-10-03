watch:
	watchman -- trigger . build '*.*' --  make build notify

stop:
	watchman watch-del .

build:
	go build -o rewl

notify:
	/usr/bin/osascript -e "display notification \"build done\""

