default:
	@just --list

alias s := start
[group("dev")]
start:
	CI=true CLICOLOR_FORCE=1 \
	DEV=1 PORT=1234 go tool air \
	-proxy.enabled=true \
	-proxy.app_port=1234 \
	-proxy.proxy_port=8080 \
	-build.delay=10 \
	-build.include_ext go,html,css,scss,png,jpg,gif,svg \
	-build.exclude_dir cache,cmd,tmp

alias u := update
# git pull and docker compose up
[group("server")]
update:
	git pull
	docker compose up -d --build
