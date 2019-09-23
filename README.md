<h1 align="center">Welcome to LabAC 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
</p>

> 展示实验室论文成果

基于gin, vue, docker, docker-compose构建网站服务。网站效果图请[点击这里](#效果图)查看。

> 本网站根据老板的需求编写, 来展示实验室的论文成果，无奈网上没有开源的代码，只能2天快速撸出这一简单版本，毕竟研究生也不是专业搞前端开发的，UI就以简洁为主，如果大家关注度比较高的话(多多star)，我会考虑继续完善这个项目(欢迎大家多提issue)。

好的README会附上一个目录

## 目录

1. [部署指南](#部署指南)
2. [迁移指南](#迁移指南)
3. [开发指南](#开发指南)
4. [技术栈](#技术栈)
5. [TODO](#TODO)
6. [效果图](#效果图)

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

      > 编译所生成的可执行文件会依赖一些库，并且是动态链接。在这里因为使用的是 `scratch`镜像，它是空镜像，因此我们需要将生成的可执行文件静态链接所依赖的库。                              --[煎鱼的博客](https://eddycjy.gitbook.io/golang/di-3-ke-gin/golang-docker#san-zhong-xin-gou-jian-jing-xiang)

   2. 进入`labac-front`目录

      ```
      npm install
      npm run build
      ```

4. 如果电脑上没有配置相关的环境请到[release](https://github.com/yangsoon/LabAC/releases)页面下载相关文件,把解压后的`labac`放在`labac-gin`文件夹内 ,`dist`文件放在`labac-front`文件夹内。

5. 在启动服务之前，有几个配置项需要配置一下

   1. `labac/conf/app.ini`文件, 修改`JWT_SECRET`和`[admin]`下的值，`JWT_SECRET`是生成token的密钥请随机输入字母和数字的组合(这里我用了我的学号),而`[admin]`下的值就是登录网站时使用的用户名和密码。

      ```ini
      RUN_MODE = debug
      JWT_SECRET = zy1806311 
      ...
      [admin]
      USERNAME = admin
      PASSWORD = admin
      ```

   2. `docker-compose.yml`,docker-compose启动时,如果端口80被占用,可以修改为其他端口, 该选项一般不需要修改。除非出现执行第6步之后出现因为端口被占用导致的无法启动的问题。

      ```yml
      version: "3"
      services:
        labac-front:
          image: "nginx:latest"
          volumes:
            - ./labac-front/dist:/usr/share/nginx/html
            - ./compose-conf/nginx.conf:/etc/nginx/conf/nginx.conf
          ports:
            - "80:80"
      ...
      ```

6. 启动服务,在项目根目录下执行:

   ```sh
   docker-compose up -d
   ```

7. 放进大象，开玩笑的。这时候输入ip就能直接访问了。

## 迁移指南

如果遇到了机器故障，需要更换机器部署应用，但是又要保证数据一致性的话，请参考下面操作。操作的前提是原本机器上的文件没有损坏，需要保证项目根目录下的`resource`和`db`文件完好，这两个文件分别存储着论文PPT资源和redis数据库可持久化数据。

1. 在新机器上下载项目代码
2. 将原本机器上的`resource`和`db`文件拷贝，并覆盖新机器上项目对应的文件。
3. 查看 **[部署指南](#部署指南)** 部署

## 开发指南



## 技术栈

本项目前端代码使用[Element](https://element.eleme.cn/#/zh-CN)-基于 Vue 2.0 的桌面端组件库构建，后端使用golang编写的web框架[gin](https://gin-gonic.com/)提供restful api服务, 使用redis存储必要的数据。整个服务使用docker-compose部署。

1. 前端 vue+vuex+vue-router+axios+element

   既然用vue写前端了，不管用到用不到都要上vue全家桶啦，前端没有什么想说的，基于一个理念就行，简单就是美，其实色彩越单调，结构越简单越好。开发的时候，突然发现加入一些好看的图标，能够极大的美化页面，因为element自带的icon种类太少，我就从阿里的[icon库](https://www.iconfont.cn/)里面选择了一些比较好的icon。发现效果确实不错，大家以后也可以使用这上面的图标。

2. 后端 golang(gin) 之前一直用python写后端，这次第一次试水了golang,发现写起来并没有很难受，感觉gin写后端很容易上手，并不比flask或者tornado差。而且golang的goroutine和channel简直不要太爽(这里没有用到，只是单纯的想夸一下golang)。学习的时候，参考的[这篇博客](https://eddycjy.gitbook.io/golang/di-3-ke-gin),写的很不错，之前golang一直被诟病没有好的包管理机制，这次使用了go mod,感觉比以前要设置path好多了。

3. docker-compose(nginx+golang+redis)

   因为老板想让部署和迁移都比较方便，那就毫无疑问上docker啦，把一些数据映射到宿主机，迁移的时候，把容器映射到宿主机的文件拷贝出来再部署就可以恢复数据了。

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