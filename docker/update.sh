#!/bin/bash
cd "$( dirname "${BASH_SOURCE[0]}" )"

for dir in */ ; do
    cd $dir
    sed -i -e "s|go get github.com/packagefoundation/yap.*|go get github.com/packagefoundation/yap # `date`|g" Dockerfile
    sudo docker build --rm -t yap/${dir::-1} .
    sed -i -e "s|go get github.com/packagefoundation/yap.*|go get github.com/packagefoundation/yap|g" Dockerfile
    cd ..
done
