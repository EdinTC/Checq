#!/bin/bash
ssh -i ../deploy_rsa $1@$2

cd checq
git fetch && git pull
docker-compose pull
docker-compose down
docker-compose up -d