#!/bin/bash
cd docker-images-pkgs||return 1
for f in *.tar
do
  echo "$f"
  set -x
  docker load -i "$f"
  set +x
done