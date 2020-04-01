#!/bin/bash
DELAY="3"
export DELAY
### 1. Install chaincode on all peers
# install cc in peer0.org1.example.com:7051
set -x
peer chaincode install \
  -n chatcc \
  -v 1.0 \
  -p chat_chaincode/
res=$?
set +x
# install cc in peer1.org1.example.com:8051
export CORE_PEER_ADDRESS=peer1.org1.example.com:8051
set -x
peer chaincode install \
  -n chatcc \
  -v 1.0 \
  -p chat_chaincode/
res=$?
set +x
# switch to Org2
export CORE_PEER_LOCALMSPID=Org2MSP
export PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp

# install cc in peer0.org2.example.com:9051
export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
set -x
peer chaincode install \
  -n chatcc \
  -v 1.0 \
  -p chat_chaincode/
res=$?
set +x
# install cc in peer1.org2.example.com:10051
export CORE_PEER_ADDRESS=peer1.org2.example.com:10051
set -x
peer chaincode install \
  -n chatcc \
  -v 1.0 \
  -p chat_chaincode/
res=$?
set +x
### 2. Instantiate the chaincode on the channel
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
set -x
peer chaincode instantiate \
  -o orderer.example.com:7050 \
  --tls \
  --cafile $ORDERER_CA \
  -C mychannel \
  -n chatcc \
  -v 1.0 \
  -c '{"Args":["init"]}' \
  -P "OR('Org1MSP.member','Org2MSP.member')" \
  --collections-config  $GOPATH/src/chat_chaincode/collections_config.json
res=$?
sleep "$DELAY"
set +x
### 3. Store private data
# Acting as a member of Org1
# Swich to peer0.org1
export CORE_PEER_LOCALMSPID=Org1MSP
export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

export CORE_PEER_ADDRESS=peer0.org1.example.com:7051

## 3.1 Invoke
# generate greebear001's public/private key
MEMBER=$(echo -n "{\"name\":\"greebear001\"}" | base64 | tr -d \\n)
export MEMBER
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["initMember"]}'  \
  --transient "{\"member\":\"$MEMBER\"}"
res=$?
set +x
# generate greebear002's public/private key
MEMBER=$(echo -n "{\"name\":\"greebear002\"}" | base64 | tr -d \\n)
export MEMBER
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["initMember"]}'  \
  --transient "{\"member\":\"$MEMBER\"}"
res=$?
set +x
# generate greebear003's public/private key
MEMBER=$(echo -n "{\"name\":\"greebear003\"}" | base64 | tr -d \\n)
export MEMBER
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["initMember"]}'  \
  --transient "{\"member\":\"$MEMBER\"}"
res=$?
sleep "$DELAY"
set +x

### 3.2 Query
# Query private data as an authorized peer
set -x
peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["queryMember","greebear001"]}'
res=$?

peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["queryMember","greebear002"]}'
res=$?
set +x

### 3.3 Save Message without Encryption
# greebear001 receive 天青色等烟雨 from greebear002
MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear002\",\"context\":\"天青色等烟雨\"}" | base64 | tr -d \\n)
export MESSAGE
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["saveMessage"]}'  \
  --transient "{\"message\":\"$MESSAGE\"}"
res=$?
set +x
# greebear001 receive 天青色等烟雨 from greebear003
MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear003\",\"context\":\"而我在等你\"}" | base64 | tr -d \\n)
export MESSAGE
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["saveMessage"]}'  \
  --transient "{\"message\":\"$MESSAGE\"}"
res=$?
set +x
# Query - Query message by receiver
set -xx
peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["queryMessagesByReceiver","greebear001"]}'
res=$?
set +x

### 3.5  Save Message with Encryption
# greebear001 receive 天青色等烟雨 from greebear002
MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear002\",\"context\":\"天青色等烟雨\"}" | base64 | tr -d \\n)
export MESSAGE
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["saveMessageUsePubKey"]}'  \
  --transient "{\"message\":\"$MESSAGE\"}"
res=$?
set +x
# greebear001 receive 而我在等你 from greebear003
MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear003\",\"context\":\"而我在等你\"}" | base64 | tr -d \\n)
export MESSAGE
set -x
peer chaincode invoke \
  -o orderer.example.com:7050 \
  --tls \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["saveMessageUsePubKey"]}'  \
  --transient "{\"message\":\"$MESSAGE\"}"
res=$?
set +x
sleep "$DELAY"
# Query - Query message by receiver with privateKey

PRIVATEKEY=$(echo -n "-----BEGIN ECC PRIVATE KEY-----\nMHcCAQEEIIKgpZ0V6k6HbWupC7+SfzB7NbxOo5YQqerKyOiur9xzoAoGCCqGSM49\nAwEHoUQDQgAEdTvKJAxd9BVlM2W9PjDKxjGThXZOhWbsfjyFhld84xbiCGTwPZwj\nGg8At0/d3pKuhmJWNPRHws3zft3oI/xtoQ==\n-----END ECC PRIVATE KEY-----")
set -x
peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c "{\"Args\":[\"queryMessagesByReceiverUsePriKey\",\"greebear001\", \"-----BEGIN ECC PRIVATE KEY-----\nMHcCAQEEIGBa1FXcrnvnVpxkw0sksL8Pvdcv7k9PuyS8w1N7oTwpoAoGCCqGSM49\nAwEHoUQDQgAEmvlwCknhEAnXirmX7pbimD5HxmpT2KpdJCXP9MQWxex20UTFYj/O\nBWBYHj9wA3tUfWpasajSlH5QmXYjsViwFw==\n-----END ECC PRIVATE KEY-----\n\"]}"
set +x