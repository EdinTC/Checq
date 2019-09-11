#!/bin/bash
cd checq
git fetch && git pull
docker-compose pull
docker-compose down
docker-compose up -d