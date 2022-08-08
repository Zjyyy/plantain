# 此Dockerfile用来构建一个Ubuntu容器，该容器结合
# VSCode 搭建 Windows下Linux-go的开发环境
FROM ubuntu:latest
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list & sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get clean & apt-get update -y
RUN apt-get install wget gcc -y
RUN wget -c https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin
RUN mkdir -p /home/dev
WORKDIR /home/dev
CMD [ "go env -w GOPROXY=https://goproxy.cn,direct" ]

# 运行起来之后执行
# go env -w GOPROXY=https://goproxy.cn,direct
# export PATH=$PATH:/usr/local/go/bin
# source ~/.bashrc

# docker build -t home/go-dev-ubuntu:1.0 .
# docker run -itd --name go-dev -p 6280:6280 -v /E/Development/plantain:/home/dev home/go-dev-ubuntu:1.0
# docker run -itd --name go-dev -p 6280:6280 --network devnet --network-alias go-dev -v <$pwd>:/home/dev home/go-dev-ubuntu:1.0