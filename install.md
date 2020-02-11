- 参考链接： https://hyperledger-fabric.readthedocs.io/en/release-1.4/tutorial/commercial_paper.html#prerequisites
## 1. 安装node.js
node.js下载网址：https://nodejs.org/en/  
Choose LTS version of node.  
解压：
```
xz -d node-v12.14.1-linux-x64.tar.xz
tar xvf node-v12.14.1-linux-x64.tar
sudo mv node-v12.14.1-linux-x64 /opt
```
(`xz -d xxx.tar.xz` 将 xxx.tar.xz解压成 xxx.tar 然后，再用 `tar xvf xxx.tar`)  
配置环境变量
```
sudo vim /etc/profile
```
加入
```
export NODE_HOME=/opt/node-v12.14.1-linux-x64
export PATH=$NODE_HOME/bin:$PATH
```
激活
```
source /etc/profile
```
PATH路径检查
```
echo $PATH
```
检查
```
node -v
npm -v
```
显示分别为
```
v12.14.1
6.13.4
```
## 2. Docker安装
### 2.1 安装
参考链接：  
https://www.runoob.com/docker/ubuntu-docker-install.html
https://www.cnblogs.com/daner1257/p/10197855.html  

修改apt源   
(备注：每次更新sources.list后必须通过 `sudo apt-get update 来更新替换`)  
(为了提高apt源速度，将源改为国内的 cn.archive.ubuntu.com ，现在 cn.archive.ubuntu.com 指向阿里云的开源镜像站 mirrors.aliyun.com)  

备份
```
sudo mv /etc/apt/sources.list /etc/apt/sources.list_backup
```
新建
```
sudo vim /etc/apt/sources.list
```
输入保持并退出
```
deb-src http://archive.ubuntu.com/ubuntu xenial main restricted #Added by software-properties
deb http://mirrors.aliyun.com/ubuntu/ xenial main restricted
deb-src http://mirrors.aliyun.com/ubuntu/ xenial main restricted multiverse universe #Added by software-properties
deb http://mirrors.aliyun.com/ubuntu/ xenial-updates main restricted
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-updates main restricted multiverse universe #Added by software-properties
deb http://mirrors.aliyun.com/ubuntu/ xenial universe
deb http://mirrors.aliyun.com/ubuntu/ xenial-updates universe
deb http://mirrors.aliyun.com/ubuntu/ xenial multiverse
deb http://mirrors.aliyun.com/ubuntu/ xenial-updates multiverse
deb http://mirrors.aliyun.com/ubuntu/ xenial-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-backports main restricted universe multiverse #Added by software-properties
deb http://archive.canonical.com/ubuntu xenial partner
deb-src http://archive.canonical.com/ubuntu xenial partner
deb http://mirrors.aliyun.com/ubuntu/ xenial-security main restricted
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-security main restricted multiverse universe #Added by software-properties
deb http://mirrors.aliyun.com/ubuntu/ xenial-security universe
deb http://mirrors.aliyun.com/ubuntu/ xenial-security multiverse
```

更新apt包索引
```
sudo apt-get update
``` 
更新apt依赖包
```
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
```
添加 Docker 的 GPG 密钥：
```
// 阿里 墙裂推荐 快到飞起
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
// 官方
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88 通过搜索指纹的后8个字符，验证您现在是否拥有带有指纹的密钥
```
sudo apt-key fingerprint 0EBFCD88
```
设置稳定版仓库
```
// 阿里 墙裂推荐 快到飞起
sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
// 官方
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```
更新 apt 包索引
```
sudo apt-get update
```
安装最新版本的 Docker Engine-Community 和 containerd
```
sudo apt-get install docker-ce docker-ce-cli containerd.io
```
测试
```
sudo docker run hello-world
```
修改docker源
```
sudo vi /etc/docker/daemon.json
```
粘贴保存
```
{"registry-mirrors": ["https://registry.docker-cn.com","http://hub-mirror.c.163.com","http://f1361db2.m.daocloud.io"]}
```
```
// 在跑fabcar实例时，用阿里镜像速度更快些
{"registry-mirrors": ["https://f1z25q5p.mirror.aliyuncs.com"]}
```
重启docker
```
sudo systemctl restart docker
```
检查
```
sudo docker info
```
### 2.2 配置非root用户加入docker用户组省去sudo权限
参考链接：https://www.cnblogs.com/caidingyu/p/10576194.html

将ubuntu用户添加到docker分组里
```
sudo usermod -aG docker ubuntu
```
检查
```
cat /etc/group | grep ubuntu
```
重启docker
```
sudo systemctl restart docker
```
需要重启终端  
原来执行`docker info`会权限不够  
现在执行`docker info`就免去了`sudo`前缀  

其他安装，后面跑脚本时需要
```
sudo apt install docker-compose
```

## 3. 安装Go
参考链接：https://golang.org/dl/
```
tar -C ~/ -xzf go1.13.6.linux-amd64.tar.gz 
sudo vim /etc/profile
// 末尾加入:
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/opt/gopath
// 激活
source /etc/profile
```
验证
```
sudo mkdir /opt/gopath/
chmod -R 777 /opt/gopath/
cd /opt/gopath/
mkdir -p ./src/hello/
vi /opt/gopath/src/hello/hello.go
```
输入并保存
```
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```
编译运行
```
cd /opt/gopath/src/hello/
go build
./hello
```

## 4. 安装Samples, Binaries 和 Docker Images
参考链接：https://raw.githubusercontent.com/hyperledger/fabric/master/scripts/bootstrap.sh
### 4.1 Samples
```
git clone -b master https://github.com/hyperledger/fabric-samples.git && cd fabric-samples && git checkout v2.0.0
```
### 4.2 Binaries
```
cp -r hyperledger-fabric-ca-linux-amd64-1.4.4.tar.gz  ~/fabric-samples
cp -r hyperledger-fabric-linux-amd64-2.0.0.tar.gz ~/fabric-samples
tar xvzf hyperledger-fabric-ca-linux-amd64-1.4.4.tar.gz
tar xvzf hyperledger-fabric-linux-amd64-2.0.0.tar.gz
```
### 4.3 Docker Images
```
set:
DOCKER=true
SAMPLES=false
BINARIES=false
```
```
chmod +x bootstrap.sh
./bootstrap.sh
```

## 5. 其他依赖安装
```
sudo apt-get install build-essential --fix-missing
```

