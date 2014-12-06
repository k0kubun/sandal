#!/usr/bin/env bash

function green() {
  printf "\e[32m"
  echo "$1"
  printf "\e[0m"
}

function red() {
  printf "\e[31m"
  echo "$1"
  printf "\e[0m"
}

pushd lang/parsing > /dev/null
make clean
make
if [ $? -ne 0 ]; then
  red "FAILED: make"
  exit 1
else
  green "PASS"
  green "ok      make"
fi
popd > /dev/null

go build
if [ $? -ne 0 ]; then
  red "FAILED: build"
  exit 1
else
  green "PASS"
  green "ok      build"
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
          red "FAILED: ${actual_filename}"
          red "    ${err_output}"
          exit 1
        else
          diff_output=`diff -u $expect_filename $tempfile`

          if [ $? -ne 0 ]; then
            red "FAILED: ${actual_filename}"
            red "${diff_output}" | awk '{ print "    " $0 }'
            exit 1
          else
            green "PASS"
            green "ok      ${actual_filename} -> ${expect_filename}"
          fi
        fi
      else
        err_output=`./sandal ${actual_filename} 2>&1 1>$tempfile`

        if [ $? -ne 0 ]; then
          red "FAILED: ${actual_filename}"
          red "    ${err_output}"
          exit 1
        else
          green "PASS"
          green "ok      ${actual_filename}"
        fi
      fi
  esac
done
