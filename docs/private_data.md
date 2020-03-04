## Private Data
### 1. Start the network
    ./byfn.sh down
    ./byfn.sh up -c mychannel -s couchdb
### 2. Install chaincode on all peers
#### 2.1 enter the CLI container
    docker exec -it cli bash
    
#### 2.2 install the Marbles chaincode 
now in Org1

peer0.org1.example.com:7051

    peer chaincode install \
    -n marblesp \
    -v 1.0 \
    -p github.com/chaincode/marbles02_private/go/

peer1.org1.example.com:8051

    export CORE_PEER_ADDRESS=peer1.org1.example.com:8051
    
    peer chaincode install \
    -n marblesp \
    -v 1.0\
     -p github.com/chaincode/marbles02_private/go/
     
switch to Org2

    export CORE_PEER_LOCALMSPID=Org2MSP
    export PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp

peer0.org2.example.com:9051

    export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
    
    peer chaincode install \
    -n marblesp \
    -v 1.0 \
    -p github.com/chaincode/marbles02_private/go/
    
peer1.org2.example.com:10051

    export CORE_PEER_ADDRESS=peer1.org2.example.com:10051
    
    peer chaincode install \
    -n marblesp \ 
    -v 1.0 cd\ 
    -p github.com/chaincode/marbles02_private/go/
    
### 3. Instantiate the chaincode on the channel
instantiate the marbles private data chaincode on the BYFN channel mychannel.

    export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    
    peer chaincode instantiate \
    -o orderer.example.com:7050 \
    --tls \
    --cafile $ORDERER_CA \
    -C mychannel \
    -n marblesp \
    -v 1.0 \
    -c '{"Args":["init"]}' \
    -P "OR('Org1MSP.member','Org2MSP.member')" \
    --collections-config  $GOPATH/src/github.com/chaincode/marbles02_private/collections_config.json

### 4. Store private data
Acting as a member of Org1
- authorized to transact with all of the private data

Swich to Org1 Peer0

    export CORE_PEER_LOCALMSPID=Org1MSP
    export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    
    export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
    
Invoke

    export MARBLE=$(echo -n "{\"name\":\"marble1\",\"color\":\"blue\",\"size\":35,\"owner\":\"tom\",\"price\":99}" | base64 | tr -d \\n)
    peer chaincode invoke \
    -o orderer.example.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["initMarble"]}'  \
    --transient "{\"marble\":\"$MARBLE\"}"
        
### 5. Query
Org1, Org2
- name, color, size, owner

Org1
- price

#### 5.1 Query private data as an authorized peer
    peer chaincode query \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["readMarble","marble1"]}'

#### 5.2 Query private data on Org2 peer 
Switch to Org2, Peer0
 
    export CORE_PEER_LOCALMSPID=Org2MSP
    export PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp

    export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
    
Query private data Org2 is authorized to

    peer chaincode query \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["readMarble","marble1"]}'

Query private data Org2 is not authorized to

    peer chaincode query \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["readMarblePrivateDetails","marble1"]}'

### 6. Purge Private Data
Switch to Org1 Peer0

    export CORE_PEER_LOCALMSPID=Org1MSP
    export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    
    export CORE_PEER_ADDRESS=peer0.org1.example.com:7051

Open a new terminal to check the log

    docker logs peer0.org1.example.com 2>&1 | grep -i -a -E 'private|pvt|privdata'

Query marble1

    peer chaincode query \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["readMarblePrivateDetails","marble1"]}'
    
Create a new marble2; Invoke  
Check the log after issue the following command

    export MARBLE=$(echo -n "{\"name\":\"marble2\",\"color\":\"blue\",\"size\":35,\"owner\":\"tom\",\"price\":99}" | base64 | tr -d \\n)
    peer chaincode invoke \
    -o orderer.example.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["initMarble"]}' \
    --transient "{\"marble\":\"$MARBLE\"}"


    export MARBLE_OWNER=$(echo -n "{\"name\":\"marble2\",\"owner\":\"joe\"}" | base64 | tr -d \\n)
    peer chaincode invoke \
    -o orderer.example.com:7050 \
    --tls \
    --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["transferMarble"]}' \
    --transient "{\"marble_owner\":\"$MARBLE_OWNER\"}"
    
Query marble1 but Purge

    peer chaincode query \
    -C mychannel \
    -n marblesp \
    -c '{"Args":["readMarblePrivateDetails","marble1"]}'