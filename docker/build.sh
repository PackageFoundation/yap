#!/bin/bash
set -e
cd "$( dirname "${BASH_SOURCE[0]}" )"

for dir in */ ; do
    cd $dir
    sudo docker build --rm -t packagefoundation/yap-${dir::-1} .
    cd ..
done
