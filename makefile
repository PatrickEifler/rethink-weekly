watch:
	watchman -- trigger . build '*.*' --  make build notify

unwatch:
	watchman watch-del .

build:
	go build -o rewl

notify:
	/usr/bin/osascript -e "display notification \"build done\""

