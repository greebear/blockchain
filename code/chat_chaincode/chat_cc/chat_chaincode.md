# Chat Chaincode
## 1. Start the network
    ./byfn.sh down
    ./byfn.sh up -c mychannel -s couchdb
## 2. Install chaincode on all peers
### 2.1 enter the CLI container
    docker exec -it cli bash
    
### 2.2 install the Chatcc chaincode
 
    export CORE_PEER_LOCALMSPID=Org1MSP
    export PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    
    export CORE_PEER_ADDRESS=peer1.org1.example.com:7051
    peer chaincode install \
      -n chatcc \
      -v 1.0 \
      -p chat_chaincode/
      
    export CORE_PEER_ADDRESS=peer1.org1.example.com:8051
    peer chaincode install \
      -n chatcc \
      -v 1.0 \
      -p chat_chaincode/
       
    export CORE_PEER_LOCALMSPID=Org2MSP
    export PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    
    export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
    
    peer chaincode install \
      -n chatcc \
      -v 1.0 \
      -p chat_chaincode/
      
    export CORE_PEER_ADDRESS=peer1.org2.example.com:10051
    
    peer chaincode install \
      -n chatcc \
      -v 1.0 \
      -p chat_chaincode/    
      
    
## 3. Instantiate the chaincode on the channel
instantiate the marbles private data chaincode on the BYFN channel mychannel.

    export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
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

## 4. Store private data
Acting as a member of Org1
### 4.1 Members
init members
- greebear001
- greebear002
- greebear003

---    
    MEMBER=$(echo -n "{\"name\":\"greebear001\"}" | base64 | tr -d \\n)
    export MEMBER
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["initMember"]}'  \
      --transient "{\"member\":\"$MEMBER\"}"
      
      
    MEMBER=$(echo -n "{\"name\":\"greebear002\"}" | base64 | tr -d \\n)
    export MEMBER
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["initMember"]}'  \
      --transient "{\"member\":\"$MEMBER\"}"
  
  
    MEMBER=$(echo -n "{\"name\":\"greebear003\"}" | base64 | tr -d \\n)
    export MEMBER
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["initMember"]}'  \
      --transient "{\"member\":\"$MEMBER\"}"

query members

    peer chaincode query \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["queryMember","greebear001"]}'
    
    
    peer chaincode query \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["queryMember","greebear002"]}'
      
    
    peer chaincode query \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["queryMember","greebear003"]}'   
      
### 4.2 noEncryptedMessages
save message  
- greebear001 receive "天青色等烟雨" from greebear002 
- greebear001 receive "而我在等你"   from greebear003

---
    MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear002\",\"context\":\"天青色等烟雨\"}" | base64 | tr -d \\n)
    export MESSAGE
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["saveMessage"]}'  \
      --transient "{\"message\":\"$MESSAGE\"}"
      
    MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear003\",\"context\":\"而我在等你\"}" | base64 | tr -d \\n)
    export MESSAGE
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["saveMessage"]}'  \
      --transient "{\"message\":\"$MESSAGE\"}"

query message

    peer chaincode query \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["queryMessagesByReceiver","greebear001"]}'  
      
          
### 4.3 EncrpytedMessages
save message
- greebear001 receive 天青色等烟雨 from greebear002
- greebear001 receive 而我在等你 from greebear003


    MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear002\",\"context\":\"天青色等烟雨\"}" | base64 | tr -d \\n)
    export MESSAGE
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["saveMessageUsePubKey"]}'  \
      --transient "{\"message\":\"$MESSAGE\"}"
    
    
    
    MESSAGE=$(echo -n "{\"receiver\":\"greebear001\",\"sender\":\"greebear003\",\"context\":\"而我在等你\"}" | base64 | tr -d \\n)
    export MESSAGE
    peer chaincode invoke \
      -o orderer.example.com:7050 \
      --tls \
      --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
      -C mychannel \
      -n chatcc \
      -c '{"Args":["saveMessageUsePubKey"]}'  \
      --transient "{\"message\":\"$MESSAGE\"}"


query message

    PRIVATEKEY="-----BEGIN ECC PRIVATE KEY-----\nMHcCAQEEIP5ndjdD+WWDB0K/lQ08sqqu2jfH74o46iGz5S6PJKmHoAoGCCqGSM49\nAwEHoUQDQgAEDzrp9/WrAu7rEojHyynSwaEW3U4hW0TDkxSXClFPnwrcDng+iaPm\nzAE14M404Uaabdjnu0EPJ7REqFByiwBQvw==\n-----END ECC PRIVATE KEY-----\n"
    peer chaincode query \
      -C mychannel \
      -n chatcc \
      -c "{\"Args\":[\"queryMessagesByReceiverUsePriKey\",\"greebear001\", \"$PRIVATEKEY\"]}"