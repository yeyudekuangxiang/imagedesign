#!/bin/bash
exist=`docker ps -a -q -f=name=imagedesign`
if [ -n "${exist}" ]; then
    docker rm -f imagedesign > /dev/null
fi
cd ../ && docker build -t imagedesign/imagedesign:latest .
docker run --name=imagedesign -d -p 80:80 imagedesign/imagedesign:latest