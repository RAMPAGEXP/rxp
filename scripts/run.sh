#!/bin/bash

set -e

GIT_TAG=$(git describe --tags)

echo "> Running $GIT_TAG..."

docker run --rm -it -p 26657:26657 --name rxp-local rampagexp/rxp:$GIT_TAG /bin/sh