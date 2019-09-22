<h1 align="center">Welcome to LabAC 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
</p>

> 展示实验室论文成果

## Install

因为golang官方镜像太大，后端的docker容器选择基于Scratch构建，所以先在本地编译好后端程序。

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o labac .
```

## Author

👤 **yangsoon**

* Github: [@yangsoon](https://github.com/yangsoon)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_