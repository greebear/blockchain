<a name="Table of Contents" />

- [私有数据](#Private-data)
- [私有数据集(PDC)定义](#PDC-definition)


<a name="Private-datas" />

## Private data

### 1. 如何保持共享私有数据
#### 1.1 比较蠢的方法:
当同一个通道下面的部分组织，希望使一些私有的数据其他组织无法访问，那么需要在这一部分组织之间创建一个新的通道。
- 缺点: 
    - 额外的管理开销：维持链码版本、背书策略、MSP(Membership Service Provider)
    - 灵活性差：无法实现加入通道的组织能查看一部分数据的同时，灵活地屏蔽一些私有数据
#### 1.2 新办法:
fabric-v1.2提供了一个称为私有数据集(private data collection)的方法:

    允许同一个通道下的部分组织，可以背书、提交、查询私有数据。
    不需要额外创建新的通道。
### 2  什么是私有数据集(private data collection)
"私有数据集"(private data collection)包含两个元素:
- 实际的私有数据
- 该数据的哈希编码

#### 2.1 实际的私有数据(The actual private data)
- 通过gossip协议点对点(peer-to-peer)传递
    - 只有被授权的组织才能看到这些数据
- 存放在私有状态数据库(private state database)
    - 该私有状态数据库也称SideDB
    - 该私有状态数据库在被授权的peer(Authorized Peer)上
    - 该私有状态数据库可被链码访问存取
- ordering service无法查看私有数据
- 需要在通道上建立锚节点并为每个锚节点配置CORE_PEER_GOSSIP_EXTERNALENDPOINT
    - 为了实现跨组织的交流

#### 2.2 该数据的哈希编码(A hash of that data)
- 此编码可以被 背书(endorsed)、排序(ordered)和写入(written)通道里每个peer的账本上
- 作为交易的凭证、状态的验证、并可用于审核目的

#### 2.3 实例
https://hyperledger-fabric.readthedocs.io/en/release-1.4/_images/PrivateDataConcept-2.png

#### 2.4 其他
- 成员可将私有数据共享给第三方查看
    - 第三方如何验证该私有数据是否正确：
        - 通过得到的私有数据进行哈希运算得到哈希编码
        - 将哈希编码与账本里通道状态下之前保存的哈希编码进行对比
            
                补充：因此必须保证账本通道状态下保存的哈希码不可被篡改
                
### 3. 使用一个通道下的私有数据集 or 使用额外的通道？   
- 私有数据集(collections)
    - 账本(或交易记录)需要在组织之间共享，但是部分组织需要能够访问一次交易中的部分(或全部)数据
        - 所以在同一个通道下，组织间交易时产生的哈希码是所有人可见的(channel state);  
          而实际的数据部分，只有授权的组织才能查看(private data state)     
- 通道(channels)    
    - 账本(或交易记录)对于组织中的部分成员需要全部保密
        - 所以separate channel外的组织无法查看任何关于交易的信息，包括哈希码
        
### 4. PDC案例总结
- 一个组织可以有多个SideDB
    - 比如区块链chat，一人(一个组织)需要与多人对话，建立多个SideDB
        - 需要考虑的问题：SideDB扩容     
  
  
-
     
<a name="PDC-definition" />
            
##  私有数据集(PDC)定义
- 当进行链码实例化的时候，PDC的定义就被部署到通道上
### 1. 构建集定义(collection definition)json文件
集定义(collection definition)功能
- 谁可以持有数据
- 将数据分布式存储在多少个节点上
- 分发私有数据需要多少个节点
- 数据库持有数据的时长

集定义(collection definition)属性
- name
- policy 
    - 定义组织中哪些节点可以持有数据集
- requiredPeerCount
    - 分发私有数据需要多少个节点，才能被链码认可
- maxPeerCount
    - 分布式储存的节点个数(背书节点)
    - 一些背书节点失效，其他背书节点可以分发私有数据
- blockToLive
    - 数据在块上存有的时间，常用于敏感数据
    - 设置为0时，用不清除数据
- memberOnlyRead
    - true 代表集的成员组织才有权限获取私有数据
    - 这里的成员组织指(policy)所写成员
    
    
    // collections_config.json
    
    [
      {
           "name": "collectionMarbles",
           "policy": "OR('Org1MSP.member', 'Org2MSP.member')",
           "requiredPeerCount": 0,
           "maxPeerCount": 3,
           "blockToLive":1000000,
           "memberOnlyRead": true
      },
    
      {
           "name": "collectionMarblePrivateDetails",
           "policy": "OR('Org1MSP.member')",
           "requiredPeerCount": 0,
           "maxPeerCount": 3,
           "blockToLive":3,
           "memberOnlyRead": true
      }
    ]
