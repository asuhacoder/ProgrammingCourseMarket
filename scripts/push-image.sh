#!/bin/bash

# cd from root dir
cd ./deployments/terraform

PRODUCT_NAME=skhole
REGION=`terraform output region`
REGION=`sed -e 's/^"//' -e 's/"$//' <<<"$REGION"`
ECR_USER_URL=`terraform output ecr-user-url`
ECR_USER_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_USER_URL"`
ECR_AUTH_URL=`terraform output ecr-auth-url`
ECR_AUTH_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_AUTH_URL"`
ECR_COURSE_URL=`terraform output ecr-course-url`
ECR_COURSE_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_COURSE_URL"`
ECR_LESSON_URL=`terraform output ecr-lesson-url`
ECR_LESSON_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$REGIONECR_LESSON_URL"`
ECR_TEST_CASE_URL=`terraform output ecr-test-case-url`
ECR_TEST_CASE_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_TEST_CASE_URL"`
ECR_GATEWAY_URL=`terraform output ecr-gateway-url`
ECR_GATEWAY_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_GATEWAY_URL"`
ECR_RUNNER_URL=`terraform output ecr-runner-url`
ECR_RUNNER_URL=`sed -e 's/^"//' -e 's/"$//' <<<"$ECR_RUNNER_URL"`

# back to root dir
cd ../../

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_USER_URL
docker build . -t $PRODUCT_NAME-user -f ./build/package/docker/user/Dockerfile
docker tag $PRODUCT_NAME-user:latest $ECR_USER_URL
docker push $ECR_USER_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_AUTH_URL
docker build . -t $PRODUCT_NAME-auth -f ./build/package/docker/auth/Dockerfile
docker tag $PRODUCT_NAME-auth:latest $ECR_AUTH_URL
docker push $ECR_AUTH_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_COURSE_URL
docker build . -t $PRODUCT_NAME-course -f ./build/package/docker/course/Dockerfile
docker tag $PRODUCT_NAME-course:latest $ECR_COURSE_URL
docker push $ECR_COURSE_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_LESSON_URL
docker build . -t $PRODUCT_NAME-lesson -f ./build/package/docker/lesson/Dockerfile
docker tag $PRODUCT_NAME-lesson:latest $ECR_LESSON_URL
docker push $ECR_LESSON_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_TEST_CASE_URL
docker build . -t $PRODUCT_NAME-test-case -f ./build/package/docker/test_case/Dockerfile
docker tag $PRODUCT_NAME-test-case:latest $ECR_TEST_CASE_URL
docker push $ECR_TEST_CASE_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_GATEWAY_URL
docker build . -t $PRODUCT_NAME-gateway -f ./build/package/docker/gateway/Dockerfile
docker tag $PRODUCT_NAME-gateway:latest $ECR_GATEWAY_URL
docker push $ECR_GATEWAY_URL

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ECR_RUNNER_URL
docker build . -t $PRODUCT_NAME-runner -f ./build/package/docker/runner/Dockerfile
docker tag $PRODUCT_NAME-runner:latest $ECR_RUNNER_URL
docker push $ECR_RUNNER_URL