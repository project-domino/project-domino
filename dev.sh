#!/bin/bash

# Exit immediately on error.
set -e;

# Client-watching function.
function watchClient() {
	cd assets;
	if [[ ! -e node_modules ]]; then
		npm i;
	fi;
	gulp watch;
};

# Server-watching function.
function watchServer() {
	command -v CompileDaemon >/dev/null 2>&1 || go get -u github.com/githubnemo/CompileDaemon;
	CompileDaemon \
		-build "go build ./..." \
		-command "cmd/project-domino-server/project-domino-server -dev" \
		-exclude-dir node_modules \
		-exclude-dir .git \
		-graceful-kill \
		-pattern "(.+\.go|.+\.c|.+\.pug)$";
};

# Ensure that a port is defined.
export PORT="${1-12345}";

(watchClient) & (watchServer);
