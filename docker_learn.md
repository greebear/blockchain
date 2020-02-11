杀死所有正在运行的容器
```
docker kill $(docker ps -a -q)
```
删除所有已经停止的容器
```
docker rm $(docker ps -a -q)
```