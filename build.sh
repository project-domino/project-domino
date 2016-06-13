#!/bin/bash

# Exit immediately on error.
set -e;

# Client-building function.
function client() {
	cd assets;
	if [[ ! -e node_modules ]]; then
		npm i;
	else
		gulp clean;
		gulp;
	fi;
};

# Server-building function.
function server() {
	go get .;
	go build;
};

# Deployment bundle-building function.
function deployBundle() {
	tar czf project-domino.tgz project-domino -C assets/dist assets.zip;
}

# Realpath function for OS X people.
# Credit https://stackoverflow.com/questions/3572105.
command -v realpath >/dev/null 2>&1 || function realpath() {
	[[ "${1}" = /* ]] && echo "${1}" || echo "$(pwd)/${1#./}";
};

# A simple function that builds, always from the base directory.
base_dir="$(dirname $(realpath ${0}))";
function build() {
	cd "${base_dir}";
	${@};
};

build client;
build server;
build deployBundle;
