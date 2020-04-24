#!/bin/bash
docker network create --subnet 192.168.28.0/24 styx
cp ../bin/styxnode .
cp ../tools/certificates/*.pem .
cp ../tools/certificates/host_key .
n=3
while [ $n -gt 0 ]
  do cp ../configs/dockerconfig"${n}".yml .
    docker stop styx"${n}"
    let n-=1
done
#docker rmi styx
docker build -t styx .
rm -v dockerconfig*
rm -v host_key
#rm -v styxmaster
rm -v styxnode
n=3
while [ $n -gt 0 ]
 do docker run -dit --net styx --ip 192.168.28.1"${n}" -e configfile=/opt/dockerconfig"${n}".yml --rm --name styx"${n}" styx
    let n-=1
done
echo "done"