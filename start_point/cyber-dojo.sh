#!/bin/bash

# There is no internet access so do not contact any module proxy
# to download or resolve modules.
export GOPROXY=off
# There is no internet access so do not verify any module's checksum
# against the public checksum database (sum.golang.org)
export GONOSUMDB='*'
# Use cache pre-populated in the docker image we are running in.
export GOCACHE=/go/build-cache

go test
