## 什么是区块链
区块链就是一个具有共享状态的密码性安全交易的单机(cryptographically secure transactional singleton machine with shared-state)。[1]这有点长，是吧？让我们将它分开来看：

- “密码性安全(Cryptographically secure)”是指用一个很难被解开的复杂数学机制算法来保证数字货币生产的安全性。将它想象成类似于防火墙的这种。它们使得欺骗系统近乎是一个不可能的事情（比如：构造一笔假的交易，消除一笔交易等等）。
- “交易的单机(Transactional singleton machine)”是指只有一个权威的机器实例为系统中产生的交易负责任。换句话说，只有一个全球真相是大家所相信的。
- “具有共享状态(With shared-state)”是指在这台机器上存储的状态是共享的，对每个人都是开放的。

## Certificate Authorities 证书颁发机构
- 发行证书给管理员和网络节点
- 所发证书: X.509
- Hyperledger Fabric提供了内置的Fabric-CA
- CA的作用:
    - 1 认证每个组件(component)从属于某个特定的组织
        - 由上可得，一个网络中不同组织会有不同的CA
    - 2 客户端的交易请求、智能合约的交易响应
        - 是交易生成和验证过程的核心