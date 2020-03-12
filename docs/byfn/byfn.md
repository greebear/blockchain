# 1 Crypto Generator 
## 1.1 generateCerts 生成证书
cryptogen generate    
- 输入: 
    - `crypto-config.yaml`
- 输出: 
    - crypto-config文件夹
        - MSP material
            - certs
            - keys
- 命令如下:

        ../bin/cryptogen generate \
        --config=./crypto-config.yaml
    
## 2 generateChannelArtifacts
### 2.1 创建Orderer创世区块
configtxgen
- 环境变量

        export FABRIC_CFG_PATH=$PWD
        
    作用：帮助configtxgen找到`configtx.yaml`文件
    
- 输入:
    - -profile
        - CONSENSUS_TYPE
            - solo: TwoOrgsOrdererGenesis
            - kafka: SampleDevModeKafka
            - etcdraft: SampleMultiNodeEtcdRaft
    - -channelID
        - byfn-sys-channel
    - -outputBlock
        - ./channel-artifacts/genesis.block
        
- 输出:
    - ./channel-artifacts/genesis.block
    
- 命令如下:

        ../bin/configtxgen \
        -profile SampleMultiNodeEtcdRaft \
        -channelID byfn-sys-channel \
        -outputBlock ./channel-artifacts/genesis.block
    
    
### 2.2 创建通道配置事务
#### 2.2.1 创建 channel transaction artifact 
【如果是】Raft或者Kafka则忽略此步  直接进入2.2.2

configtxgen
- 环境变量
    
        export CHANNEL_NAME=mychannel

- 输入        
    - -profile
        - CONSENSUS_TYPE
            - solo: TwoOrgsOrdererGenesis
            - kafka: SampleDevModeKafka
            - etcdraft: SampleMultiNodeEtcdRaft
    - -channelID
        - $CHANNEL_NAME
    - -outputCreateChannelTx
        - ./channel-artifacts/channel.tx

- 输出
    - ./channel-artifacts/channel.tx

- 命令如下:  

        ../bin/configtxgen \
        -profile  TwoOrgsChannel \
        -outputCreateChannelTx ./channel-artifacts/channel.tx \
        -channelID $CHANNEL_NAME 

 ### 2.2.2 定义 anchor peer
 configtxgen
 
 - 环境变量
 
        export CHANNEL_NAME=mychannel
- 输入:                   
    - -profile
        - CONSENSUS_TYPE
            - solo: TwoOrgsOrdererGenesis
            - kafka: SampleDevModeKafka
            - etcdraft: SampleMultiNodeEtcdRaft    
    - -channelID
        - $CHANNEL_NAME
    - -asOrg 
        - Org1MSP
    - -outputAnchorPeersUpdate
        - ./channel-artifacts/Org1MSPanchors.tx
- 输出:
    - ./channel-artifacts/Org1MSPanchors.tx         
    
- 命令如下:

        ../bin/configtxgen \
        -profile TwoOrgsChannel \
        -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx \
        -channelID $CHANNEL_NAME \
        -asOrg Org1MSP

-

        ../bin/configtxgen \
        -profile TwoOrgsChannel \
        -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx \
        -channelID $CHANNEL_NAME \
        -asOrg Org2MSP


           