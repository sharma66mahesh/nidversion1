#!/bin/bash
USER_NAME="ubuntu"
cd /home/${USER_NAME}/nid-version-1
sudo chown -R ${USER_NAME} ./web
sudo chmod +x nidnetwork/startNetwork.sh
sudo chmod +x web/startFabric.sh