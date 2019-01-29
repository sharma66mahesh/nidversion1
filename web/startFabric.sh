#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

CHANNEL_NAME="nid-channel"

#check if chaincode name is passed
if [ $# == 0 ]; then
	CHAINCODE_NAME="nidchain2"
	CHAINCODE_VERSION="2.0"
elif [ $# == 2 ]; then
	CHAINCODE_NAME="$1"
	CHAINCODE_VERSION="$2"
fi

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
LANGUAGE=${1:-"golang"}
CC_SRC_PATH=github.com/chaincode/enrollmentChaincode
if [ "$LANGUAGE" = "node" -o "$LANGUAGE" = "NODE" ]; then
	CC_SRC_PATH=/opt/gopath/src/github.com/fabcar/node
fi

# clean the keystore
rm -rf ./hfc-key-store

# launch network; create channel and join peer to channel
cd ../nidnetwork
./startNetwork.sh

# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars
# docker-compose -f ./docker-compose.yml up -d cli

#check if chaincode is already installed
OUTPUT=$(docker exec cli peer chaincode list --installed | grep $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec -e "CORE_PEER_LOCALMSPID=mohaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/crypto/peerOrganizations/moha.nid.com/users/Admin@moha.nid.com/msp" cli peer chaincode install -n $CHAINCODE_NAME -v $CHAINCODE_VERSION -p "$CC_SRC_PATH" -l "$LANGUAGE"	
fi

#check if chaincode is already instantiated on channel
OUTPUT=$(docker exec cli peer chaincode list --instantiated -C $CHANNEL_NAME | grep $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec -e "CORE_PEER_LOCALMSPID=mohaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/crypto/peerOrganizations/moha.nid.com/users/Admin@moha.nid.com/msp" cli peer chaincode instantiate -o orderer.nid.com:7050 -C $CHANNEL_NAME -n $CHAINCODE_NAME -l "$LANGUAGE" -v $CHAINCODE_VERSION -c '{"Args":[""]}'
fi
sleep 1
# docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n main -c '{"function":"initLedger","Args":[""]}'

printf "\nTotal setup execution time : $(($(date +%s) - starttime)) secs ...\n\n\n"
printf "Start by installing required packages run 'npm install'\n"
printf "Then run 'node enrollAdmin.js', then 'node registerUser'\n\n"
printf "The 'node invoke.js' will fail until it has been updated with valid arguments\n"
printf "The 'node query.js' may be run at anytime once the user has been registered\n\n"
