#!/bin/bash

USER_NAME="ubuntu"
CHANNEL_NAME="nid-channel"
CHAINCODE_NAME="nidchain"
CHAINCODE_VERSION="4.0"
#check if chaincode name is passed
# if [ $# == 0 ]; then
# 	CHAINCODE_NAME="nidchain2"
# 	CHAINCODE_VERSION="2.0"
# elif [ $# == 2 ]; then
# 	CHAINCODE_NAME="$1"
# 	CHAINCODE_VERSION="$2"
# else
# 	printf "Error: Invalid args\n"
# fi

# don't rewrite paths for Windows Git Bash users
starttime=$(date +%s)
CC_SRC_PATH=github.com/chaincode/kailashChaincode

# launch network; create channel and join peer to channel
cd /home/${USER_NAME}/nid-version-1/nidnetwork

rm -rf ../web/hfc-key-store

if [ "$1" == "prune" ]; then
	./startNetwork.sh prune
else 
	./startNetwork.sh
fi


#check if chaincode is already installed
OUTPUT=$(docker exec cli peer chaincode list --installed | grep -w $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec cli peer chaincode install -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -p "$CC_SRC_PATH"
fi
sleep 1
#check if chaincode is already instantiated on channel
OUTPUT=$(docker exec cli peer chaincode list --instantiated -C "$CHANNEL_NAME" | grep -w $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec cli peer chaincode instantiate -o orderer.nid.com:7050 -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -c '{"Args":[""]}'
fi
sleep 10
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n main -c '{"function":"initLedger","Args":[""]}'

printf "\nTotal setup execution time : $(($(date +%s) - starttime)) secs ...\n\n\n"
printf "Start by installing required packages run 'npm install'\n"
printf "Then run 'node enrollAdmin.js', then 'node registerUser'\n\n"
printf "The 'node invoke.js' will fail until it has been updated with valid arguments\n"
printf "The 'node query.js' may be run at anytime once the user has been registered\n\n"
