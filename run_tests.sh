#!/usr/bin/env bash

pushd lang/parsing > /dev/null
make clean
make
if [ $? -ne 0 ]; then
  printf "\e[31m"
  echo "FAILED: make"
  printf "\e[0m"
  exit 1
else
  printf "\e[32m"
  echo "PASS"
  echo "ok      make"
  printf "\e[0m"
fi
popd > /dev/null

go build
if [ $? -ne 0 ]; then
  printf "\e[31m"
  echo "FAILED: build"
  printf "\e[0m"
  exit 1
else
  printf "\e[32m"
  echo "PASS"
  echo "ok      build"
  printf "\e[0m"
fi

testdirs=$(find . -type f -name "*\_test.go" | grep -v "\.gondler" | sed -e "s/[^\/]*$//" | uniq)

for testdir in $testdirs; do
  pushd $testdir > /dev/null
  if result=$(go test); then
    printf "\e[32m${result}\e[0m\n"
  else
    printf "\e[31m${result}\e[0m\n"
    exit 1
  fi
  popd > /dev/null
done

tempfile=`mktemp /tmp/tmp.XXXXXXXXXX`
trap "rm $tempfile" 0

for filename in `find test -type f -name '*.sandal'`; do
  case $filename in
    (*.sandal)
      actual_filename="${filename}"
      expect_filename="${filename%.sandal}.smv"

      if [ -e $expect_filename ]; then
        err_output=`./sandal ${actual_filename} 2>&1 1>$tempfile`

        if [ $? -ne 0 ]; then
          printf "\e[31m"
          echo "FAILED: ${actual_filename}"
          echo "    ${err_output}"
          printf "\e[0m"
          exit 1
        else
          diff_output=`diff -u $expect_filename $tempfile`

          if [ $? -ne 0 ]; then
            printf "\e[31m"
            echo "FAILED: ${actual_filename}"
            echo "${diff_output}" | awk '{ print "    " $0 }'
            printf "\e[0m"
            exit 1
          else
            printf "\e[32m"
            echo "PASS"
            echo "ok      ${actual_filename} -> ${expect_filename}"
            printf "\e[0m"
          fi
        fi
      else
        err_output=`./sandal ${actual_filename} 2>&1 1>$tempfile`

        if [ $? -ne 0 ]; then
          printf "\e[31m"
          echo "FAILED: ${actual_filename}"
          echo "    ${err_output}"
          printf "\e[0m"
          exit 1
        else
          printf "\e[32m"
          echo "PASS"
          echo "ok      ${actual_filename}"
          printf "\e[0m"
        fi
      fi
  esac
done
