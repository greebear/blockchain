## Certificate Authorities 证书颁发机构
- 发行证书给管理员和网络节点
- 所发证书: X.509
- Hyperledger Fabric提供了内置的Fabric-CA
- CA的作用:
    - 1 认证每个组件(component)从属于某个特定的组织
        - 由上可得，一个网络中不同组织会有不同的CA
    - 2 客户端的交易请求、智能合约的交易响应
        - 是交易生成和验证过程的核心
        - 