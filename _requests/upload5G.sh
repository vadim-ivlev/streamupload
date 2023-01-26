#!/bin/bash
echo 'Uploading 5G file'

dir=$(dirname "$0")

file=$dir/5G

curl -vvv -H "Content-Type:application/octet-stream" -T $file -X POST localhost:19090/streamupload