```shell
go run vartest_main.go #运行go进程,vartest_main.go里面必须是main包,要有main方法
go build vartest_main.go #编译vartest_main.go到可执行文件

```

```shell
go mod init 包名 #包名一般是和目录名一样,也可以不一样
go mod tidy #拉取缺少的模块，移除不用的模块
go mod vendor #将依赖复制到vendor下
go run vartest_main.go vartest1.go #一般只有main包,只需要run一个文件就行
go env -w GOPROXY=https://goproxy.cn,direct #设置全局代理
```