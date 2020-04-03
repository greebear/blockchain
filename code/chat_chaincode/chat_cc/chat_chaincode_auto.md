## chat chaincode
### 1. Start the network
    ./byfn.sh down
    ./byfn.sh up -c mychannel -s couchdb
### 2. Auto Test
#### 2.1 enter the CLI container
    docker exec -it cli bash

#### 2.2 run the auto script    
    $GOPATH/src/chat_chaincode/chat_cc_build.sh