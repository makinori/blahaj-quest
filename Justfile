default:
	@just --list

[private]
generate:
	go tool templ generate

alias r := run
run: generate
	DEV=1 go run .

alias w := watch
alias s := watch
alias start := watch
watch:
	DEV=1 go tool templ generate \
	--watch --proxy="http://127.0.0.1:8080" \
	--cmd "go run ."

# #!/usr/bin/env sh
# echo "ctrl+c to restart, hold it down to close"
# while true; do
# 	just run
# 	sleep 0.1
# done

alias b := build
build: generate
	CGO_ENABLED=0 GOOS=linux go build -o blahaj-quest
	strip blahaj-quest
