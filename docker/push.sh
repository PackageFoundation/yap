#!/bin/bash
set -e
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

for distro in $(ls -d) ; do
    docker push "packagefoundation/yap-${distro}:latest"
done
