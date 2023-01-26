#!/bin/bash
echo 'Creating files for testing'

dir=$(dirname "$0")

file=$dir/100M
dd if=/dev/zero of=$file bs=1M count=100 status=progress
file=$dir/1G
dd if=/dev/zero of=$file bs=1M count=1024 status=progress
file=$dir/5G
dd if=/dev/zero of=$file bs=1M count=5170 status=progress
