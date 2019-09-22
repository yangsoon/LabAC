<h1 align="center">Welcome to LabAC 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
</p>

> 展示实验室论文成果

基于gin, element, docker, docker-compose构建的容器化网站。网站效果图请[点击这里](#效果图)查看。

> 本网站根据老板的需求编写, 来展示实验室的论文成果，无奈网上没有开源的代码，只能疯狂coding 2,3天写出了这一简单版本，毕竟研究生做前端开发的也不多，UI就以简洁为主，如果大家关注度比较高的话(多多star)，那么欢迎大家多提issue,我会考虑继续完善这个项目。

## 部署指南

1. 首先在要部署的机器上安装docker和docker-compose, docker的安装请根据机器系统参考网上的安装教程，docker-compose的安装建议使用 pip安装 (如果docker-compose下载较慢输入下面的代码,切换到豆瓣镜像)。

   ```shell
   pip install -i http://pypi.douban.com/simple/ --trusted-host=pypi.douban.com/simple docker-compose
   ```

2. 从仓库下载源码

   ```
   git clone https://github.com/yangsoon/LabAC.git
   ```

3. 如果机器上已经有了golang1.12和nodejs的环境,请继续阅读,否则查看步骤4。

   因为golang官方docker镜像太大，为了减轻镜像下载时网络和机器硬盘空间的压力，后端服务的docker镜像选择基于Scratch构建，所以在启动服务之前请先在本地编译好后端代码。

   同时前端代码是使用vue-cli脚手架编写，build之后的文件没有放在仓库里面。所以如果有node环境，请先自己build。请参考 **[开发指南](#开发指南)**。

   具体步骤如下(前提是机器上已经有了golang1.12和nodejs的环境):

   1. 进入`labac-gin`目录，输入下面指令，会得到`labac`的可执行文件

      ```sh
      CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o labac .
      ```

   2. 进入`labac-front`目录

      ```
      npm install
      npm run build
      ```

4. 如果电脑上没有配置相关的环境请到[release](https://github.com/yangsoon/LabAC/releases)页面下载相关文件,把解压后的`labac`放在`labac-gin`文件夹内 ,`dist`文件放在`labac-front`文件夹内。

5. 在启动服务之前，有几个配置项需要配置一下

   1. `labac/conf/app.ini`文件, 修改`JWT_SECRET`和`[admin]`下的值，`JWT_SECRET`是生成token的密钥,而`[admin]`下的值就是登录网站时使用的用户名和密码。

      ```ini
      RUN_MODE = debug
      JWT_SECRET = zy1806311 
      ...
      [admin]
      USERNAME = admin
      PASSWORD = admin
      ```

   2. config.js 下面修改后端请求地址，为机器的ip

   3. `docker-compose.yml`,docker-compose启动时,如果端口50被占用,可以修改为其他端口, 该选项一般不需要修改。除非出现执行第5步之后出现因为端口被占用导致的无法启动的问题。

      ```yml
      version: "3"
      services:
        labac-front:
          image: "nginx:latest"
          volumes:
            - ./labac-front/dist:/usr/share/nginx/html
            - ./compose-conf/nginx.conf:/etc/nginx/conf/nginx.conf
          ports:
            - "50:80"
      ```

6. 启动服务,在项目根目录下执行:

   ```sh
   docker-compose up -d
   ```

## 迁移指南

如果遇到了机器故障，需要更换机器部署应用，但是又要保证数据一致性的话，请参考下面操作。操作的前提是原本机器上的文件没有损坏，需要保证项目根目录下的`resource`和`db`文件完好，这两个文件分别存储着论文PPT资源和redis数据库可持久化数据。

1. 在新机器上下载项目代码
2. 将原本机器上的`resource`和`db`文件拷贝，并覆盖新机器上项目对应的文件。
3. 查看 **[部署指南](#部署指南)** 部署

## 开发指南

本项目前端代码使用[Element](https://element.eleme.cn/#/zh-CN)-基于 Vue 2.0 的桌面端组件库构建，后端使用golang编写的web框架[gin](https://gin-gonic.com/)提供restful api服务, 使用redis存储必要的数据。整个服务使用docker-compose部署。

## TODO

* [ ] 添加首页UI和首页可编辑功能

## 效果图



## Author

👤 **yangsoon**

* Github: [@yangsoon](https://github.com/yangsoon)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_