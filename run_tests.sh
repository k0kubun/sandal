#!/usr/bin/env bash

testdirs=$(find . -type f -name "*\_test.go" | grep -v "\.gondler" | sed -e "s/[^\/]*$//" | uniq)

for testdir in $testdirs; do
  pushd $testdir > /dev/null
  if result=$(go test); then
    printf "\e[32m${result}\e[0m\n"
  else
    printf "\e[31m${result}\e[0m\n"
  fi
  popd > /dev/null
done
