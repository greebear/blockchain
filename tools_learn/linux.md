 ##### 1. 查看挂载路径 内存占用量 df的意思是dish file
    df -h

##### 2. 查看当前各目录大小
    du -sh *

#### 3. 挂载硬盘 思路：创建分区、创建文件系统、挂载 
sdb(spatial database)  
对于经常使用的设备可写入文件/etc/fastab,以使系统在每次开机时自动加载

#### 4. 查看cpu核数
    cat /proc/cpuinf

#### 5. 环境变量相关：
https://www.cnblogs.com/youyoui/p/10680329.html

#### 6. 创建自定义命令
语法: `alias[别名]=[指令名称]`  
eg.  

    vim ~/.bash_profile
    // 加入
    alias sshbc1="cd someplace; vagrant ssh blockchain-1"
    // 输入下面命令不用重启terminal
    source ~/.bash_profile 
    
#### 7. 查看网络通不通
- ping  
    仅针对域名或ip
- telnet  
    可域名(ip) +
- wget
- nslookup    
