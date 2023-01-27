#!/bin/bash
echo 'upload 100M file'

dir=$(dirname "$0")

file=$dir/100M

curl -vvv -H "Content-Type: multipart/form-data" -F "file=@$file" -X POST localhost:19090/postupload