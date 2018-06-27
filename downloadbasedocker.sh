#!/bin/bash -eu
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#


##################################################
# This script pulls docker images from hyperledger
# docker hub repository and Tag it as
# hyperledger/fabric-<image> latest tag
##################################################

#Set ARCH variable i.e ppc64le,s390x,x86_64,i386
ARCH=`uname -m`

BASE_DOCKER_TAG="x86_64-0.4.6"
DOCKER_NS="hyperledger"
REGISTERY_URL="docker.m-chain.com"

dockerFabricPull() {
    docker pull $REGISTERY_URL/$DOCKER_NS/mchain-baseos:$BASE_DOCKER_TAG
    docker tag $REGISTERY_URL/$DOCKER_NS/mchain-baseos:$BASE_DOCKER_TAG $DOCKER_NS/mchain-baseos
    docker pull $REGISTERY_URL/$DOCKER_NS/mchain-baseimage:$BASE_DOCKER_TAG
    docker tag $REGISTERY_URL/$DOCKER_NS/mchain-baseimage:$BASE_DOCKER_TAG $DOCKER_NS/mchain-baseimage
    docker pull $REGISTERY_URL/$DOCKER_NS/mchain-couchdb:$BASE_DOCKER_TAG
    docker tag $REGISTERY_URL/$DOCKER_NS/mchain-couchdb:$BASE_DOCKER_TAG $DOCKER_NS/mchain-couchdb
    docker pull $REGISTERY_URL/$DOCKER_NS/mchain-zookeeper:$BASE_DOCKER_TAG
    docker tag $REGISTERY_URL/$DOCKER_NS/mchain-zookeeper:$BASE_DOCKER_TAG $DOCKER_NS/mchain-zookeeper
    docker pull $REGISTERY_URL/$DOCKER_NS/mchain-kafka:$BASE_DOCKER_TAG
    docker tag $REGISTERY_URL/$DOCKER_NS/mchain-kafka:$BASE_DOCKER_TAG $DOCKER_NS/mchain-kafka
}

echo "===> Pulling base Images"
dockerFabricPull

echo
echo "===> List out hyperledger docker images"
docker images | grep hyperledger*
