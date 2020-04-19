#!/bin/bash
docker network create --subnet 192.168.28.0/24 styx
cp ../bin/styxnode .
cp ../configs/ssh_host_rsa_key .
n=3
while [ $n -gt 0 ]
  do cp ../configs/dockerconfig"${n}".yml .; let n-=1
done
docker build -t styx .
rm -v dockerconfig*
rm -v ssh_host_rsa_key
#rm -v styxmaster
rm -v styxnode
n=3
while [ $n -gt 0 ]
 do docker run -dit --net styx --ip 192.168.28.1"${n}"  --rm --name styx"${n}" styx
    docker exec -dit styx"${n}" /styxnode --config dockerconfig"${n}".yml --overwrite-hostconfig
    let n-=1
done