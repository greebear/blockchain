# 基于Fabric的区块链聊天APP
## 1. 代码
**地址:**
https://github.com/greebear/blockchain/blob/master/code/chat_chaincode/


**结构:**
```
chat_chaincode
├── SHA1Float64.go              用于生成数字签名，但此场景用不上
├── chat_cc.go                  主函数入口
├── collections_config.json     私有数据集配置
├── eccErytion.go               ECC加密算法
├── ecies.go                    以太坊加密
├── getEccKey.go                生成ECC公钥私钥
├── membership.go               成员信息保存、访问的链码
├── messageMngm.go              消息保存、访问的链码
├── richQuery.go                信息访问辅助函数
└── chat_cc
    ├── chat_cc_build.sh        测试脚本
    ├── chat_chaincode.md       相关指令
    └── chat_chaincode_auto.md  自动测试指令
```
## 2. 聊天APP
![BlockChainChat](./assets/BlockChainChat.png) 
### 2.1 私有数据集配置
相关配置可见` collections_config.json`
#### 2.1.1 CA - collectionMembers
这里的CA指的不是fabric原生的CA，而是用于储存用户公钥、私钥的Collection。
为什么称之为“Collection”，因为fabric私有数据保存时，可以设置不同的Collection，每个Collection中可以储存不通的数据，如下图所示。
![SideDB-org1-org2](./assets/SideDB-org1-org2.png) 
这里fabric官网教程中设置了Marbles和MarblesPrivateData两个Collection，Colleciton的功能有两种：
- 1. 可以设置不同的Collection来存放不同关键字、不同类型的数据，类似于数据库中不同的表
- 2. 可以设置哪些Org有访问权限

这里应用其第一种功能——“储存不同数据”，用于生成不同的“表”来储存数据，CA则用于储存用户的公钥、私钥。
```
题外话：
其实一个更好的用户管理CA，不应将“私钥”存储于网络中，此处为了方便管理，将公私钥一并储存。
开发者可根据情况，或让用户自己生成公私钥，并修改`membership.go`的内容，仅用于保存公钥。
```
### 2.1.2 明文储存 - collectionMessages
若用户在本地已用ECC算法加密数据，上传数据到cli的话，那么直接保存数据到collectionMessages，无需进行“在线加密”。
### 2.1.3 密文储存 - collectionMessagesEnrcypt
用于在线加密储存消息

## 2.2 实例展示
### 1. 启动网络
所用的组织结构为byfn的结构，数据库为couchdb(可提供丰富的查询功能)
```bash
./byfn.sh down
./byfn.sh up -c mychannel -s couchdb
```
### 2. 在所有节点上安装链码
#### 2.1 进入CLI容器
```bash
docker exec -it cli bash
```
#### 2.2 安装Chatcc链码
```bash
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
```
### 3. 实例化Chatcc链码
```bash
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
```
### 4. 保存数据
Acting as a member of Org1
#### 4.1 Members
生成用户
- greebear001
- greebear002
- greebear003
```bash
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
```

查询用户
```bash
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
```
### 4.2 明文储存
保存明文消息，不进行在线加密
- greebear001 receive "天青色等烟雨" from greebear002
- greebear001 receive "而我在等你" from greebear003

```bash
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
```
查询消息
```bash
peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c '{"Args":["queryMessagesByReceiver","greebear001"]}'  
```

### 4.3 密文储存
保存密文消息，进行在线加密
- greebear001 receive 天青色等烟雨 from greebear002
- greebear001 receive 而我在等你 from greebear003

```bash
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
```

查询消息，这里查询的为greebear001所收到的所有消息，所以需要有greebear001的私钥才能看到明文内容。
可用前面展示的命令查询得到greebear001的私钥，替换掉`PRIVATEKEY`的内容。
```bash
PRIVATEKEY="-----BEGIN ECC PRIVATE KEY-----\nMHcCAQEEIP5ndjdD+WWDB0K/lQ08sqqu2jfH74o46iGz5S6PJKmHoAoGCCqGSM49\nAwEHoUQDQgAEDzrp9/WrAu7rEojHyynSwaEW3U4hW0TDkxSXClFPnwrcDng+iaPm\nzAE14M404Uaabdjnu0EPJ7REqFByiwBQvw==\n-----END ECC PRIVATE KEY-----\n"
peer chaincode query \
  -C mychannel \
  -n chatcc \
  -c "{\"Args\":[\"queryMessagesByReceiverUsePriKey\",\"greebear001\", \"$PRIVATEKEY\"]}"
```
## 3. 参考资料
- https://hyperledger-fabric.readthedocs.io/en/latest/private_data_tutorial.html
- https://github.com/wumansgy/goEncrypt