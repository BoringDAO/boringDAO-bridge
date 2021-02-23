#!/usr/bin/env bash

set -e

source x.sh

# $1 is arch, $2 is source code path
case $1 in
linux-amd64)
  print_blue "Compile for linux/amd64"
  docker run -t \
    -v $2:/code/boring-bridge \
    -v ~/.ssh:/root/.ssh \
    -v ~/.gitconfig:/root/.gitconfig \
    -v $GOPATH/pkg/mod:$GOPATH/pkg/mod \
    golang:1.15 \
    /bin/bash -c "go env -w GO111MODULE=on &&
      go env -w GOPROXY=https://goproxy.cn,direct &&
      go get -u github.com/gobuffalo/packr/packr &&
      cd /code/boring-bridge && make install &&
      mkdir -p /code/boring-bridge/bin &&
      cp /go/bin/bridge /code/boring-bridge/bin/bridge"
  ;;
*)
  print_red "Other architectures are not supported yet"
  ;;
esac
