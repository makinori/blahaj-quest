default:
	@just --list

[private]
generate:
	go tool templ generate

alias r := run
run: generate
	DEV=1 go run .

alias b := build
build: generate
	CGO_ENABLED=0 GOOS=linux go build -o blahaj-quest
	strip blahaj-quest

# https://templ.guide/developer-tools/live-reload-with-other-tools/

[private]
watch-main:
	DEV=1 go tool templ generate \
	--watch --proxy "http://127.0.0.1:8080" \
	--open-browser=false \
	--cmd "go run ."

[private]
watch-notify:
	go tool air \
	--build.cmd "go tool templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.include_dir "ui" \
	--build.include_ext "js,css" \
	-misc.clean_on_exit "true" \
	--log.main_only "true"

alias w := watch
alias s := watch
alias start := watch
watch:
	#!/usr/bin/env bash
	set -e
	pids=()

	cleanup() {
		for pid in "${pids[@]}"; do
			kill "$pid" || true
		done
		exit 1
	}
	trap cleanup SIGINT SIGTERM

	commands=(
		"just watch-main"
		"sleep 1 && just watch-notify"
	)

	for cmd in "${commands[@]}"; do
		bash -c "$cmd" &
		pids+=($!)
	done

	wait

# #!/usr/bin/env sh
# echo "ctrl+c to restart, hold it down to close"
# while true; do
# 	just run
# 	sleep 0.1
# done

