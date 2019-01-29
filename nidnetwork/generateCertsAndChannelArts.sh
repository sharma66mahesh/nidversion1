#!/bin/bash

#generate certs
./bin/cryptogen generate --config=./crypto-config.yaml

export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=nid-channel
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/nid-channel.tx -channelID $CHANNEL_NAME

#anchor peers update transaction
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/mohaMSPanchors.tx -channelID $CHANNEL_NAME -asOrg mohaMSP
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/ecMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ecMSP