#!/bin/sh
cp ./interface-grpc/*.proto ./interface-grpc/environment-java/src/main/proto
cd ./interface-grpc/environment-java
gradle clean generateProto
cd ../../
cp -r ./interface-grpc/environment-java/generated ./client-java
