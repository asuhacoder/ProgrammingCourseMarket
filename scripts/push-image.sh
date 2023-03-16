#!/bin/bash
cd ../../
PRODUCT_NAME=skhole
docker build . -t $PRODUCT_NAME-$2 -f ./build/package/docker/$2/Dockerfile
docker tag $PRODUCT_NAME-$2:latest $1


docker push $1