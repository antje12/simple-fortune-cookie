#!/bin/bash 
echo "sleeping"
sleep 20s

RESULT=$(curl -o /dev/null -s -w "%{http_code}\n" http://34.79.255.98:32540/)

if [ $RESULT == '200' ]
then exit 0
else exit 1
fi
