#!/usr/bin/env bash

source x.sh
CURRENT_PATH=$(pwd)
PROJECT_PATH=$(dirname "${CURRENT_PATH}")
BUILD_PATH=${PROJECT_PATH}/dist
CONFIG_PATH=${PROJECT_PATH}/config
APP_VERSION=$(if [ "$(git rev-parse --abbrev-ref HEAD)" == "HEAD" ]; then git describe --tags HEAD; else echo "dev"; fi)

rm -rf "$BUILD_PATH"
mkdir -p "$BUILD_PATH"
print_blue "===> 1. Install packr"
if ! type packr >/dev/null 2>&1; then
  go get -u github.com/gobuffalo/packr/packr
fi

print_blue "===> 2. Build bridge and plugin"
cd "${PROJECT_PATH}" && make build && make plugin

print_blue "===> 3. Pack binary"
cd "${PROJECT_PATH}" || (echo "project path is not exist" && return)
cp ./bin/bridge "${BUILD_PATH}"/
cp -r ./build/*.so  "${BUILD_PATH}"/
cd "${BUILD_PATH}"
if [ "$(uname)" == "Darwin" ]; then
  tar -zcvf bridge_darwin_x86_64_"${APP_VERSION}".tar.gz "${BUILD_PATH}"/*
else
  tar -zcvf bridge_linux_x86_64_"${APP_VERSION}".tar.gz "${BUILD_PATH}"/*
fi