### 0. Vagrant安装
#### 0.1 vagrant版本问题
- Vagrant 2.2.6 doesn't work with VirtualBox 6.1.0

https://github.com/oracle/vagrant-boxes/issues/178?utm_source=hacpai.com

【issue】
The provider 'virtualbox' that was requested to back the machine
'aubuntu-1' is reporting that it isn't usable on this system. The
reason is shown below:

Vagrant has detected that you have a version of VirtualBox installed
that is not supported by this version of Vagrant. Please install one of
the supported versions listed below to use Vagrant:

4.0, 4.1, 4.2, 4.3, 5.0, 5.1, 5.2, 6.0

A Vagrant update may also be available that adds support for the version
you specified. Please check www.vagrantup.com/downloads.html to download
the latest version.
【end】

### 1. Vagrantfile 的配置 
https://www.jianshu.com/p/bdc66b8bafbc

### 2. 创建自己的box
#### 2.1 相关命令
查看虚拟机列表
  
    vboxmanage list vms

查看全局运行状态

    vagrant global-status
打包
    
    vagrant package --base blockchain-1 --output ./fabric-prerequisites
    

- base 要打包的虚拟机名称
- output 打包后的包名
- include 打包需要增加的文件，多个文件以逗号分隔
- vagrantfile 指定vagrantfile文件  

#### 2.2 实例 ubuntu16.04出现ssh问题的解决方案【一】
参考于：https://github.com/hashicorp/vagrant/issues/5186
##### 2.2.1 启动虚拟环境
启动

    vagrant up
    
如果出现需要输入密码的情况，则输入

    vi ~/.vagrant.d/boxes/xenial/0/virtualbox/Vagrantfile
    
可见如下，复制对应密码输入即可

    config.ssh.username = "ubuntu"
    config.ssh.password = "6e673d6bcae167481cxxxxxx" 
    
##### 2.2.2 修改密码
修改ubuntu用户的密码，比如"vagrant"

    sudo passwd ubuntu
    
##### 2.2.3 打包及box添加

    vagrant package --base blockchain-1 --output ./fabric.box
    vagrant box add fabric ./fabric.box
    
检查

    vagrant global-status
    
##### 2.2.4 生成虚拟环境
在init的Vagrantfile中加入
 
    config.ssh.username = 'ubuntu'
    config.ssh.password = 'vagrant'   
    
启动

    vagrant up
    
##### 2.2.5 总结
- 打包前的用户密码，需要【修改】得和打包后，`vagrant up`新的box时所加载Vagrantfile中配置的用户密码一致
- 【备注：与上面不同的其他方法】  
    尝试了将官网下载的ubuntu/xenial64.box添加到box list后，默认生成的Vagrantfile文件，替换掉新添加的`fabric.box`自动生成的Vagrantfile,w
    
        cp ~/.vagrant.d/boxes/xenial/0/virtualbox/Vagrantfile ~/.vagrant.d/boxes/fabric/0/virtualbox/
##### 2.2.2 
### 3. vagrant 常见命令

- vagrant up 启动
- vagrant halt 关机
- vagrant reload 重启
- vagrant destroy 移除机器
- vagrant ssh 登陆

box管理命令

- vagrant box list 查看本地box列表
- vagrant box add 添加box到列表
- vagrant box remove 从box列表移除

