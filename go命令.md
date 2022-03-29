```shell
go mod init 包名 #包名一般是和目录名一样,也可以不一样
go mod tidy #拉取缺少的模块，移除不用的模块
go mod vendor #将依赖复制到vendor下
go run vartest_main.go vartest1.go #一般只有main包,只需要run一个文件就行
```