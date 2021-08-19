#!/bin/bash
[[ -z "${GIT_COMMIT}" ]] && Tag='local' || Tag="${GIT_COMMIT::4}"
[[ -z "${docker_username}" ]] && DockerRepo='' || DockerRepo="${docker_username}/"
[[ -z "$1" ]] && lol1='' || lol1="$1"

echo "This is a thing: ${lol1}"
docker build -t "${DockerRepo}$ENV_NAME:latest" -t "${DockerRepo}$ENV_NAME:1.0-$Tag" ./"$ENV_NAME"