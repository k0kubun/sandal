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

tempfile=`mktemp /tmp/tmp.XXXXXXXXXX`
trap "rm $tempfile" 0

for filename in `ls test`; do
  case $filename in
    (*.sandal)
      actual_filename="test/${filename}"
      expect_filename="test/${filename%.sandal}.expect"

      if [ -e $expect_filename ]; then
        err_output=`./sandal ${actual_filename} 2>&1 1>$tempfile`

        if [ $? -ne 0 ]; then
          printf "\e[31m"
          echo "FAILED: ${actual_filename}"
          echo "    ${err_output}"
          printf "\e[0m"
        else
          diff_output=`diff -u $expect_filename $tempfile`

          if [ $? -ne 0 ]; then
            printf "\e[31m"
            echo "FAILED: ${actual_filename}"
            echo "${diff_output}" | awk '{ print "    " $0 }'
            printf "\e[0m"
          else
            printf "\e[32m"
            echo "PASSED: ${actual_filename}"
            printf "\e[0m"
          fi
        fi
      fi
  esac
done
