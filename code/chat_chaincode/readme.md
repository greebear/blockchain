
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

