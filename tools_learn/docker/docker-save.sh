#!/bin/bash
mkdir docker-images-pkgs
cd docker-images-pkgs||return 1
for i in $(docker images|grep fabric|awk '{print $1}'|uniq)
do
  set -x
  docker save -o  "${i#*/}.tar" "$i"
  set +x
done


