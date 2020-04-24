#!/bin/sh
set -ex
apt install -y unzip
mkdir protobuf-setup && cd protobuf-setup
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
unzip protoc-3.11.4-linux-x86_64.zip
cp -r include/* /usr/bin/include/
cp -r bin/* /usr/bin/

