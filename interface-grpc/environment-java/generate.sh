#!/bin/sh
cp ./interface-grpc/*.proto ./interface-grpc/environment-java/src/main/proto
cd ./interface-grpc/environment-java
gradle clean generateProto
cd ../../
# client-javaへコピー
cp -r ./interface-grpc/environment-java/build/generated/source/proto/main/java/sample ./client-java/src/main/java
cp -r ./interface-grpc/environment-java/build/generated/source/proto/main/grpc/sample ./client-java/src/main/java
