#!/bin/bash

#generate certs
./bin/cryptogen generate --config=./crypto-config.yaml

FABRIC_CFG_PATH=$PWD
CHANNEL_NAME="nid-channel"
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID "$CHANNEL_NAME"

#anchor peers update transaction
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/MohaMSPanchors.tx -channelID "$CHANNEL_NAME" -asOrg MohaMSP
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/EcMSPanchors.tx -channelID "$CHANNEL_NAME" -asOrg EcMSP