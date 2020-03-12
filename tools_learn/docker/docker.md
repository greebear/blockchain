杀死所有正在运行的容器

    docker kill $(docker ps -a -q)

删除所有已经停止的容器

    docker rm $(docker ps -a -q)
    
-    
    docker rm -f $(docker ps -aq)
    docker network prune
    docker volume prune
    

删除所有镜像

    docker rmi -f `docker images -q`

系统配置

    /lib/systemd/system/docker.service
    
docker关闭重启等

    systemctl stop docker
    
    sudo systemctl restart docker
    
docker镜像储存位置修改
- 需要`systemctl stop docker`后再`sudo systemctl restart docker`才能生效

        {
        "registry-mirrors": ["http://hub-mirror.c.163.com"],
        "graph": "/home/ubuntu/docker"
        }
