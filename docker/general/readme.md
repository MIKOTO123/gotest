# docker的go环境搭建

# 简介

该项目更适用于本机没有go环境,为了永久搭建go环境,例如你所有go项目的根目录在goproject里面,将Dockerfile文件放到goproject底下

# 运行说明

```docker
docker build -t goproject/go-env  . #构建镜像
docker run -itd  --name goenv -v "$PWD":/usr/src/goproject goproject/go-env #运行镜像
docker exec -it goenv /bin/bash #运行镜像
```

# 这里记录一个小插曲,













