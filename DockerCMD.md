# 基于Docker创建的开发环境的Docker指令汇总

### docker创建容器间通信网络
```docker
# 创建桥接网络
docker network create -d bridge devnet
```
### 基于go-dev镜像创建容器
```docker
# 最小go-dev开发容器
docker run -itd --name go-dev -p 6280:6280 -v /E/Development/plantain:/home/dev home/go-dev-ubuntu:1.0

# 带网络的容器，方便多容器通信
docker run -itd --restart always --name go-dev -p 6280:6280 -p 8080:8080 --network devnet --network-alias go-dev -v <$pwd>:/home/dev home/go-dev-ubuntu:1.0
```

### 基于influx镜像创建容器
```docker
# 最基本influxdb
docker run -itd --name influx-dev --restart always -p 8086:8086 --network devnet --network-alias influxdb-dev influxdb:latest
```