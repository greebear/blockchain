## Peers
- peer nodes简称peers
- peer是网络的基本元素
- peer管理(host)账本和智能合约
- 更准确度地说：peer实际上管理账本的实例和链码(智能合约)的实例
    - 为什么说用"实例"更准确，因为每单个peer中可以管理多种账本和链码
    - 账本永久地记录所有由智能合约生成的交易
    - 在Hyperledger Fabric中，智能合约包含在链码里
        - 有时可以简称链码为智能合约？
            - TODO:https://hyperledger-fabric.readthedocs.io/en/release-1.4/smartcontract/smartcontract.html
- peer通过链码去获取分布式账本的副本
- peer可以被创建、启动、停止、配置、删除
- 应用和管理员必须通过peer进行交互，使其能访问所有资源

### 多账本 多链码
- 1个peer管理1个或多个账本
- 1个账本有0个或多个链码应用于它
    - 虽然可以出现1个账本内0个链码的情况，但是非常少见
    - 链码允许peer去查询(query)和更新(update)账本
    - peer中，无论用户是否安装了外部应用提供的链码，peer始终存在特殊的系统链码
        - TODO:小朋友，你是否有很多问号??? 那么这个特殊的系统链码不算入链码的个数，不然就不可能出现0个链码的情况，不愧是我。
- peer上的不同账本可以拥有同一套链码
- 例子：
    - P1上有2个账本L1, L2
        - L1上有1个链码S1
        - L2上有2个链码S1, S2
       
### 应用
- Fabric Software Development Kit (SDK) 简化了这个过程
    - 使得应用(applications)可以：
        - 连接peer
        - 调用(invoke)链码去生成、提交交易到网络
        - 接收流程处理完成事件提醒
        
- Ledger-query 账本查询
    - 在应用(applications)与peer间涉及一个【三步对话】
        - 1 app连接到peer
        - 2 调用链码(invoke chaincode [proposal])
            - 2.1 peer调用链码
            - 2.2 链码生成请求响应(proposal response), 其中包含查询结果(query result)或拟定账本更新(proposed ledger update)
        - 3 peer返回请求结果给app
        
- Ledger-update 账本更新
    - 在应用(applications)与peer间涉及一个【五步对话】
        - 1-3 同上
        - 4 app向orderer发送交易排序请求
    
#### XXX
- peers和orderers一起保证每个peer上的账本都是最新的
    
