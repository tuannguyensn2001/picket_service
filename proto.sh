#!/bin/bash

cd proto

protos=`ls *.proto`
cd ../
#
for proto in $protos
do
  IFS='.' read -r name var2 <<< "$proto"
  echo "start generate $name"
  result=$?
  make gen-proto name="$name"
  if [ $result -eq 0 ]; then
    echo "generate $name success"
  fi
  if [ $result -ne 0 ]; then
      echo "generate $name fail"
    fi
done