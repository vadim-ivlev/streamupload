#!/bin/bash


. $(dirname "$0")/git_repos


git pull $gitlab master
git pull $github master

echo "Подключение к сети elastic"
nmcli con up elastic

git pull origin  master

echo "Отключение от сети elastic"
nmcli con down elastic

