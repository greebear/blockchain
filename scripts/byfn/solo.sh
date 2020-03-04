#!/usr/bin/env bash

# generate
./byfn.sh generate

# bring the network up (solo)
./byfn.sh up

# generate the certs and keys
../bin/cryptogen generate --config=./crypto-config.yaml

# create the orderer genesis block (solo)
export FABRIC_CFG_PATH=$PWD
../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block

# create the channel transaction artifac
export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

# define the anchor peer for Org1&Org2 on the channel
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

# start our network
docker-compose -f docker-compose-cli.yaml up

# enter the CLI container
docker exec -it cli bash
