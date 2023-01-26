#!/bin/bash

echo "Введите комментарий к коммиту:"
read msg
echo "Комментарий = $msg."
echo 


git add -A .
git commit -m "$msg."


. $(dirname "$0")/git_repos



git push $gitlab --all #master
git push $github --all #master
git push $gitlab --tags 
git push $github --tags 

echo "VPN ON -----------------------"
nmcli con up elastic

git push origin --all #master
git push origin --tags 

echo "VPN OFF ----------------------"
nmcli con down elastic
