#!/bin/bash
PATH=${PATH}:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/bin
USER_NAME="ubuntu"
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

cd /home/${USER_NAME}/nid-version-1/web
npm install

rm -rf /home/${USER_NAME}/nid-version-1/web/hfc-key-store
/home/${USER_NAME}/nid-version-1/web/startFabric.sh
node /home/${USER_NAME}/nid-version-1/web/enrollAdmin.js
node /home/${USER_NAME}/nid-version-1/web/registerUser.js
forever start -a -l /home/${USER_NAME}/logs.txt -o /home/${USER_NAME}/output.txt -e /home/${USER_NAME}/error.txt /home/${USER_NAME}/nidversion1/web/app.js