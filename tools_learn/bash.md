https://blog.csdn.net/s2421458535/article/details/101022029

    -r file　　　　　用户可读为真 
    -w file　　　　　用户可写为真 
    -x file　　　　　用户可执行为真 
    -f file　　　　　文件为正规文件为真 
    -d file　　　　　文件为目录为真 
    -c file　　　　　文件为字符特殊文件为真 
    -b file　　　　　文件为块特殊文件为真 
    -s file　　　　　文件大小非0时为真 
    -t file　　　　　当文件描述符(默认为1)指定的设备为终端时为真
    
    -eq           //等于
    -ne           //不等于    
    -gt           //大于
    -lt           //小于
    ge            //大于等于
    le            //小于等于
    
    $?：          获取函数返回值或者上一个命令的退出状态
    
https://www.cnblogs.com/mlfz/p/11427760.html

    -z 判断 变量的值，是否为空； zero = 0
    - 变量的值，为空，返回0，为true
    - 变量的值，非空，返回1，为false
    -n 判断变量的值，是否为空 name = 名字
    - 变量的值，为空，返回1，为false
    - 变量的值，非空，返回0，为true
    pid="123"
    [ -z "$pid" ] 单对中括号变量必须要加双引号
    [[ -z $pid ]] 双对括号，变量不用加双引号
     
    [ -n "$pid" ] 单对中括号，变量必须要加双引号
    [[ -z $pid ]] 双对中括号，变量不用加双引号    
    
bash 下 ; && || 的区别
   
    cmd1 ; cmd2	cmd1 和 cmd2 都会 被执行
    cmd1 && cmd2	如果 cmd1 执行 成功 则执行 cmd2
    cmd1 || cmd2	如果 cmd1 执行 失败 则执行 cmd2