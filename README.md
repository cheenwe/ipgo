# IP 转换

根据IP返回IP所属区域(国家省份和地市信息)

http://localhost:8081/assets/

实现效果:

![image](https://user-images.githubusercontent.com/5643208/83135569-963ba600-a118-11ea-8202-d420016ca54d.png)


## 接口示例

http://0.0.0.0:8081/ip?ip=116.21.33.1

```
{"code":1,"data":{"Ip":"116.21.33.1","Country":"中国","Province":"广东","City":"广州"},"msg":"成功"}
```

## dev

```
go run main.go
```


## build

```

GOOS=windows GOARCH=amd64  go build
GOOS=linux GOARCH=amd64  go build
GOOS=darwin GOARCH=amd64  go build
```
生成可执行文件时需要同时把 assets 目录放到可执行文件同级目录
