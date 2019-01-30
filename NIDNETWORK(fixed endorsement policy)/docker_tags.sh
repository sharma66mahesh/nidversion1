#!/bin/bash

if [ "$1" == "remove" ]; then
    
    if [ $# -eq 2 ]; then
        docker rmi hyperledger/fabric-ca:${2}
        docker rmi hyperledger/fabric-tools:${2}
        docker rmi hyperledger/fabric-peer:${2}
        docker rmi hyperledger/fabric-orderer:${2}
        docker rmi hyperledger/fabric-ccenv:${2}
    else
        printf "PASS THE TAG TO REMOVE\n"
    fi
fi

if [ "$1" == "tag" ]; then
    if [ $# -eq 3 ]; then
        docker tag hyperledger/fabric-ca:${2} hyperledger/fabric-ca:${3}
        docker tag hyperledger/fabric-tools:${2} hyperledger/fabric-tools:${3}
        docker tag hyperledger/fabric-peer:${2} hyperledger/fabric-peer:${3}
        docker tag hyperledger/fabric-ccenv:${2} hyperledger/fabric-ccenv:${3}
        docker tag hyperledger/fabric-orderer:${2} hyperledger/fabric-orderer:${3}
    fi
fi