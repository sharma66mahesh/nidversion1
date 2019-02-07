#!/bin/bash
sleep 20
docker stop $(docker ps -q)
docker rm $(docker ps -qa)
yes | docker volume prune
yes | docker network prune
CHAINCODE_IMAGES=`docker images | grep "^dev-peer" | awk '{print $3}'`

if [ "$CHAINCODE_IMAGES" != "" ]; then
        echo "Removing chaincode docker images ..."
        docker rmi -f $CHAINCODE_IMAGES
fi

cd /home/ec2-user/nid-version-1/web
npm install

rm -rf /home/ec2-user/nid-version-1/web/hfc-key-store
/home/ec2-user/nid-version-1/web/startFabric.sh
node /home/ec2-user/nid-version-1/web/enrollAdmin.js
node /home/ec2-user/nid-version-1/web/registerUser.js
forever start -a -l /home/ec2-user/logs.txt -o /home/ec2-user/output.txt -e /home/ec2-user/error.txt /home/ec2-user/nidversion1/web/app.js
