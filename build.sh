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
	cd cmd/project-domino-server;
	go get -v .;
	go build -v;

	cd cmd/project-domino-mail;
	go get -v .;
	go build -v;
};

# Deployment bundle-building function.
function deployBundle() {
	tmp_file="$(mktemp)";
	tar rf "${tmp_file}" -C cmd/project-domino-mail project-domino-mail;
	tar rf "${tmp_file}" -C cmd/project-domino-server project-domino-server;
	tar rf "${tmp_file}" -C assets/dist assets.zip;
	gzip "${tmp_file}" -c > project-domino.tgz;
	rm "${tmp_file}";
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
	echo "-----> ${@}";
	${@};
};

build client;
build server;
build deployBundle;

echo "-----> Done building!";
