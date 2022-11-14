#!/usr/bin/env bash

set -e

docker run --rm --name buildkite-agent -d -t \
    -e BUILDKITE_AGENT_TOKEN=${BUILDKITE_AGENT_TOKEN} \
    -v "/var/lib/buildkite/builds:/var/lib/buildkite/builds" \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    -v "${PWD}/.buildkite/secrets:/buildkite/secrets:ro" \
    -v "${PWD}/.buildkite/hooks:/buildkite/hooks:ro" \
    buildkite/agent:3