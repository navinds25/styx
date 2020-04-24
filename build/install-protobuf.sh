#!/bin/sh
set -ex
sudo apt install -y unzip
mkdir protobuf-setup && cd protobuf-setup
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
unzip protoc-3.11.4-linux-x86_64.zip
sudo cp -r include/* /usr/bin/include/
sudo cp -r bin/* /usr/bin/
#go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
