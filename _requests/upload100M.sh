#!/bin/bash
echo 'upload 100M file'

dir=$(dirname "$0")

file=$dir/100M

curl -vvv -H "Content-Type:application/octet-stream" -T $file -X POST localhost:19090/streamupload