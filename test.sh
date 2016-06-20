#!/bin/bash

# Exit immediately on error.
set -e;

# Realpath function for OS X people.
# Credit https://stackoverflow.com/questions/3572105.
command -v realpath >/dev/null 2>&1 || function realpath() {
	[[ "${1}" = /* ]] && echo "${1}" || echo "$(pwd)/${1#./}";
};
base_dir="$(dirname $(realpath ${0}))";

if [[ "$(pwd)" = "${base_dir}" ]]; then
	echo "-----> Currently in repo root. Moving to test subdir...";
	if [[ -d test ]]; then
		rm -r test;
	fi;
	mkdir test;
	cd test;
fi;

echo "-----> Rebuilding...";
"${base_dir}/build.sh";

echo "-----> Uncompressing...";
tar zxf "${base_dir}/project-domino.tgz";

echo "-----> Running...";
PORT=3000 ./project-domino-server;
