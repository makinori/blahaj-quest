default:
	@just --list

alias b := build
build:
	CGO_ENABLED=0 GOOS=linux go build -o blahaj-quest
	strip blahaj-quest

alias s := start
start:
	#!/usr/bin/env sh
	echo "ctrl+c to restart, hold it down to close"
	while true; do
		DEV=1 go run .
		sleep 0.1
	done
