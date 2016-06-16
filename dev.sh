#!/bin/bash

# Exit immediately on error.
set -e;

# Client-watching function.
function watchClient() {
	cd assets;
	if [[ ! -e node_modules ]]; then
		npm i;
	fi;
	gulp watch-dev;
};

# Server-watching function.
function watchServer() {
	command -v CompileDaemon >/dev/null 2>&1 || go get -u github.com/githubnemo/CompileDaemon;
	CompileDaemon \
		-command "./project-domino -dev -serveOn :${PORT}" \
		-exclude-dir node_modules \
		-graceful-kill \
		-pattern ".+(\\.c)|(\\.go)|(\\.pug)$";
};

# Ensure that a port is defined.
PORT="${1-12345}";

(watchClient) & (watchServer);
