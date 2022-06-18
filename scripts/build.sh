#!/bin/bash

set -e

GIT_TAG=$(git describe --tags)

echo "> Building $GIT_TAG..."

docker build . -t rampagexp/rxp:$GIT_TAG