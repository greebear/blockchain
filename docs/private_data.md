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

### 7. Useage
    /*
    Copyright IBM Corp. All Rights Reserved.
    
    SPDX-License-Identifier: Apache-2.0
    */
    
    // ====CHAINCODE EXECUTION SAMPLES (CLI) ==================
    
    // ==== Invoke marbles, pass private data as base64 encoded bytes in transient map ====
    //
    // export MARBLE=$(echo -n "{\"name\":\"marble1\",\"color\":\"blue\",\"size\":35,\"owner\":\"tom\",\"price\":99}" | base64 | tr -d \\n)
    // peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}"
    //
    // export MARBLE=$(echo -n "{\"name\":\"marble2\",\"color\":\"red\",\"size\":50,\"owner\":\"tom\",\"price\":102}" | base64 | tr -d \\n)
    // peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}"
    //
    // export MARBLE=$(echo -n "{\"name\":\"marble3\",\"color\":\"blue\",\"size\":70,\"owner\":\"tom\",\"price\":103}" | base64 | tr -d \\n)
    // peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}"
    //
    // export MARBLE_OWNER=$(echo -n "{\"name\":\"marble2\",\"owner\":\"jerry\"}" | base64 | tr -d \\n)
    // peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["transferMarble"]}' --transient "{\"marble_owner\":\"$MARBLE_OWNER\"}"
    //
    // export MARBLE_DELETE=$(echo -n "{\"name\":\"marble1\"}" | base64 | tr -d \\n)
    // peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["delete"]}' --transient "{\"marble_delete\":\"$MARBLE_DELETE\"}"
    
    // ==== Query marbles, since queries are not recorded on chain we don't need to hide private data in transient map ====
    // peer chaincode query -C mychannel -n marblesp -c '{"Args":["readMarble","marble1"]}'
    // peer chaincode query -C mychannel -n marblesp -c '{"Args":["readMarblePrivateDetails","marble1"]}'
    // peer chaincode query -C mychannel -n marblesp -c '{"Args":["getMarblesByRange","marble1","marble4"]}'
    //
    // Rich Query (Only supported if CouchDB is used as state database):
    //   peer chaincode query -C mychannel -n marblesp -c '{"Args":["queryMarblesByOwner","tom"]}'
    //   peer chaincode query -C mychannel -n marblesp -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'
    
    // INDEXES TO SUPPORT COUCHDB RICH QUERIES
    //
    // Indexes in CouchDB are required in order to make JSON queries efficient and are required for
    // any JSON query with a sort. As of Hyperledger Fabric 1.1, indexes may be packaged alongside
    // chaincode in a META-INF/statedb/couchdb/indexes directory. Or for indexes on private data
    // collections, in a META-INF/statedb/couchdb/collections/<collection_name>/indexes directory.
    // Each index must be defined in its own text file with extension *.json with the index
    // definition formatted in JSON following the CouchDB index JSON syntax as documented at:
    // http://docs.couchdb.org/en/2.1.1/api/database/find.html#db-index
    //
    // This marbles02_private example chaincode demonstrates a packaged index which you
    // can find in META-INF/statedb/couchdb/collection/collectionMarbles/indexes/indexOwner.json.
    // For deployment of chaincode to production environments, it is recommended
    // to define any indexes alongside chaincode so that the chaincode and supporting indexes
    // are deployed automatically as a unit, once the chaincode has been installed on a peer and
    // instantiated on a channel. See Hyperledger Fabric documentation for more details.
    //
    // If you have access to the your peer's CouchDB state database in a development environment,
    // you may want to iteratively test various indexes in support of your chaincode queries.  You
    // can use the CouchDB Fauxton interface or a command line curl utility to create and update
    // indexes. Then once you finalize an index, include the index definition alongside your
    // chaincode in the META-INF/statedb/couchdb/indexes directory or
    // META-INF/statedb/couchdb/collections/<collection_name>/indexes directory, for packaging
    // and deployment to managed environments.
    //
    // In the examples below you can find index definitions that support marbles02_private
    // chaincode queries, along with the syntax that you can use in development environments
    // to create the indexes in the CouchDB Fauxton interface.
    //
    
    //Example hostname:port configurations to access CouchDB.
    //
    //To access CouchDB docker container from within another docker container or from vagrant environments:
    // http://couchdb:5984/
    //
    //Inside couchdb docker container
    // http://127.0.0.1:5984/
    
    // Index for docType, owner.
    // Note that docType and owner fields must be prefixed with the "data" wrapper
    //
    // Index definition for use with Fauxton interface
    // {"index":{"fields":["data.docType","data.owner"]},"ddoc":"indexOwnerDoc", "name":"indexOwner","type":"json"}
    
    // Index for docType, owner, size (descending order).
    // Note that docType, owner and size fields must be prefixed with the "data" wrapper
    //
    // Index definition for use with Fauxton interface
    // {"index":{"fields":[{"data.size":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
    
    // Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
    //   peer chaincode query -C mychannel -n marblesp -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":\"marble\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'
    
    // Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
    //   peer chaincode query -C mychannel -n marblesp -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":{\"$eq\":\"marble\"},\"owner\":{\"$eq\":\"tom\"},\"size\":{\"$gt\":0}},\"fields\":[\"docType\",\"owner\",\"size\"],\"sort\":[{\"size\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'
    
