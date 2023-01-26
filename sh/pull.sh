#!/bin/bash


. $(dirname "$0")/git_repos


git pull $gitlab master
git pull $github master

echo "VPN ON -----------------------"
nmcli con up elastic

git pull origin  master

echo "VPN OFF ----------------------"
nmcli con down elastic

