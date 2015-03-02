#!/bin/bash
eval "$(gimme 1.4.1)"
export GOPATH=$HOME/gopath
export PATH=$HOME/gopath/bin:$PATH
mkdir -p $HOME/gopath/src/github.com/penlook/socket
rsync -az ${TRAVIS_BUILD_DIR}/ $HOME/gopath/src/github.com/penlook/socket/
export TRAVIS_BUILD_DIR=$HOME/gopath/src/github.com/penlook/socket
cd $HOME/gopath/src/github.com/penlook/socket