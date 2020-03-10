杀死所有正在运行的容器
```
docker kill $(docker ps -a -q)
```
删除所有已经停止的容器
```
docker rm $(docker ps -a -q)
```
```
docker rm -f $(docker ps -aq)
docker network prune
docker volume prune
```
删除所有镜像

docker rmi -f `docker images -q`