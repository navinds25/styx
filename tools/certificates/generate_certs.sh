#!/bin/bash

function certs() {
  if [ -f "ca-key.pem" ]; then
    echo "ca already exists"
  else
    cfssl gencert -initca ca-csr.json | cfssljson -bare ca
    cfssl gencert \
      -ca=ca.pem \
      -ca-key=ca-key.pem \
      -config=ca-config.json \
      -hostname=localhost,127.0.0.1,192.168.28.11,192.168.28.12,192.168.28.13 \
      -profile=styx \
      styxnode-csr.json | cfssljson -bare styxnode
  fi
}

function sshkey() {
  if [ -f "host_key" ]; then
    echo "host_key already exists"
  else
    ssh-keygen -t rsa -N "" -f host_key && rm -v host_key.pub
  fi
}

certs
sshkey

#	ssh-keygen -A -f ${CWD}
# ssh-keygen -t rsa -N "" -f host_key && rm -v host_key.pub
