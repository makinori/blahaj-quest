default:
	@just --list

alias s := start
[group("dev")]
start:
	#!/bin/bash
	set -euo pipefail

	which air &> /dev/null || {
		echo "please go install github.com/air-verse/air@latest" >&2
		exit 1
	}

	CI=true CLICOLOR_FORCE=1 \
	DEV=1 PORT=1234 air \
	-proxy.enabled=true \
	-proxy.app_port=1234 \
	-proxy.proxy_port=8080 \
	-build.delay=10 \
	-build.include_ext go,html,css,scss,js,png,jpg,gif,svg \
	-build.exclude_dir cache,cmd,tmp

alias ujs := update-js
# update javascript libraries
[group("dev")]
update-js:
	curl -Lo public/js/maplibre-gl.js https://unpkg.com/maplibre-gl@latest/dist/maplibre-gl.min.js
	curl -Lo public/css/maplibre-gl.css https://unpkg.com/maplibre-gl@latest/dist/maplibre-gl.css

alias u := update
# git pull and docker compose up
[group("server")]
update:
	git pull
	docker compose up -d --build

