#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

CHANNEL_NAME="nid-channel"
CHAINCODE_NAME="car_reg0"
CHAINCODE_VERSION="1.0"
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

CC_SRC_PATH=github.com/chaincode


# clean the keystore

# launch network; create channel and join peer to channel
# cd ../nidnetwork

if [ "$1" == "prune" ]; then
	./startNetwork.sh prune
else 
	./startNetwork
fi

# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars
# docker-compose -f ./docker-compose.yml up -d cli

#check if chaincode is already installed
OUTPUT=$(docker exec cli peer chaincode list --installed | grep -w $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec cli peer chaincode install -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -p "$CC_SRC_PATH"	
fi
sleep 2

OUTPUT=$(docker exec -e CORE_PEER_ADDRESS=peer0.ec.nid.com:7051 -e CORE_PEER_LOCALMSPID=EcMSP -e CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ec.nid.com/users/Admin@ec.nid.com/msp cli peer chaincode list --installed | grep -w $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec -e CORE_PEER_ADDRESS=peer0.ec.nid.com:7051 -e CORE_PEER_LOCALMSPID=EcMSP -e CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ec.nid.com/users/Admin@ec.nid.com/msp cli peer chaincode install -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -p "$CC_SRC_PATH"
fi
sleep 2
#check if chaincode is already instantiated on channel
OUTPUT=$(docker exec cli peer chaincode list --instantiated -C "$CHANNEL_NAME" | grep -w $CHAINCODE_NAME)
if [ -z "$OUTPUT" ]; then
	docker exec cli peer chaincode instantiate -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -c '{"Args":["init"]}' -P "AND('EcMSP.member', 'MohaMSP.member')"
fi


sleep 3
docker exec cli peer chaincode invoke -o orderer.nid.com:7050 --peerAddresses peer0.moha.nid.com:7051 --peerAddresses peer0.ec.nid.com:7051 -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["createCar","1","bmw","gtx","blue","me"]}'
sleep 2
docker exec cli peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["queryCar","1"]}'
docker exec -e CORE_PEER_ADDRESS=peer0.ec.nid.com:7051 -e CORE_PEER_LOCALMSPID=EcMSP -e CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ec.nid.com/users/Admin@ec.nid.com/msp cli peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["queryCar","1"]}'