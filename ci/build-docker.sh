#!/bin/bash
[[ -z "${GIT_COMMIT}" ]] && Tag='local' || Tag="${GIT_COMMIT::4}"
[[ -z "${docker_username}" ]] && DockerRepo='' || DockerRepo="${docker_username}/"
X=$1
echo "$X"
docker build -t "${DockerRepo}$ENV_NAME:latest" -t "${DockerRepo}$ENV_NAME:1.0-$Tag" ./"$ENV_NAME"