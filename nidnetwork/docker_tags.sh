#!/bin/bash
#usage: docker remove latest -- this removes latest tags from docker images
#usage: docker tag 1.4 latest -- this tags 1.4 version of docker images to latest

if [ "$1" == "remove" ]; then
    read -p "ENSURE YOU HAVE PASSED THE CORRECT TAGS OR YOUR IMAGES WILL BE PERMANENTLY REMOVED? (y/n)  " choice
    case $choice in    
        y)  if [ $# -eq 2 ]; then
                docker rmi hyperledger/fabric-ca:${2}
                docker rmi hyperledger/fabric-tools:${2}
                docker rmi hyperledger/fabric-peer:${2}
                docker rmi hyperledger/fabric-orderer:${2}
                docker rmi hyperledger/fabric-ccenv:${2}
            else
                printf "PASS THE TAG TO REMOVE\n"
            fi
            # ^--beginning with, $--ending with
            docker images | awk '$1~/^hyperledger/'
            ;;
        *)  echo -e "\n"
            ;;
    esac
fi

if [ "$1" == "tag" ]; then
    if [ $# -eq 3 ]; then
        docker tag hyperledger/fabric-ca:${2} hyperledger/fabric-ca:${3}
        docker tag hyperledger/fabric-tools:${2} hyperledger/fabric-tools:${3}
        docker tag hyperledger/fabric-peer:${2} hyperledger/fabric-peer:${3}
        docker tag hyperledger/fabric-ccenv:${2} hyperledger/fabric-ccenv:${3}
        docker tag hyperledger/fabric-orderer:${2} hyperledger/fabric-orderer:${3}
        docker images | grep " $3"
    fi
fi