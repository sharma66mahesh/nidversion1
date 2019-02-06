#########################USAGE#################
#create required fabric containers: 2 orgs (moha and ec) with 2 peers each. One SOLO orderer. One cli container. Ccenv (chaincode containers) are dynamically created as chaincodes are instantiated.
#pass "prune" as argument if u want to remove previously created containers, volumes and networks which might interfere with the current network. (eg:previously created channels, instantiated chaincodes,etc)
echo "Ensure you set the IMAGE_TAG environment variable to latest"
set -x #echo commands as they are typed
CHANNEL_NAME="nid-channel"

function networkUp() {
    echo "Starting the network and waiting for 10 seconds for network to fully start up"
    /usr/local/bin/docker-compose -f docker-compose-cli.yaml up -d
    sleep 10
}

function createAndJoinChannel() {
    docker exec cli peer channel create -o orderer.nid.com:7050 -c $CHANNEL_NAME -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/channel.tx
    #add peer0 from moha org
    docker exec cli peer channel join -b ${CHANNEL_NAME}.block
    #add peer1 from moha org
    #docker exec -e  CORE_PEER_ADDRESS=peer1.moha.nid.com:7051 cli peer channel join -b ${CHANNEL_NAME}.block
    #add peer from another org
    docker exec -e CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ec.nid.com/users/Admin@ec.nid.com/msp -e CORE_PEER_ADDRESS=peer0.ec.nid.com:7051 -e CORE_PEER_LOCALMSPID=EcMSP -e CORE_PEER_TLS_ENABLED=false cli peer channel join -b ${CHANNEL_NAME}.block
    #update anchor peers
    docker exec cli peer channel update -o orderer.nid.com:7050 -c $CHANNEL_NAME -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/MohaMSPanchors.tx
    docker exec -e CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ec.nid.com/users/Admin@ec.nid.com/msp -e CORE_PEER_ADDRESS=peer0.ec.nid.com:7051 -e CORE_PEER_LOCALMSPID=EcMSP -e CORE_PEER_TLS_ENABLED=false cli peer channel update -o orderer.nid.com:7050 -c $CHANNEL_NAME -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/EcMSPanchors.tx
}

echo "Stopping all running containers"
docker stop $(docker ps -q)
sleep 2

if [ "$1" == "prune" ]; then
    echo "Removing all containers, unused volumes and networks"
    docker rm $(docker ps -aq)
    docker volume prune
    docker network prune
    # Remove chaincode docker images (starting with dev-peer). awk prints 3rd column
    CHAINCODE_IMAGES=`docker images | grep "^dev-peer" | awk '{print $3}'`
    if [ "$CHAINCODE_IMAGES" != "" ]; then
        log "Removing chaincode docker images ..."
        docker rmi -f $CHAINCODE_IMAGES
    fi
fi

networkUp

#check if channel is already created. If not create and join peers to it

OUTPUT=$(docker exec cli peer channel list | grep -w $CHANNEL_NAME)

if [ -z "$OUTPUT" ]; then
    createAndJoinChannel
fi

# list all containers. Find lines with "hyperledger" on second column. Print their first column i.e. ID
#docker ps -a | awk '$2~/hyperledger/ {print $1}'
