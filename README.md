> The documentation was written on 23 May 2022.

<p align="center">
    <img src="README/logo.svg" alt="logo">
</p>
<h1 align="center">OneTiny</h1>
<div align="center">
  <a href="https://www.boii.xyz">
    <img src="https://img.shields.io/badge/BLOG-Boii.xyz-1296DB.svg?style=for-the-badge" alt="BLOG">
  </a>
</div>
<br/>
<div align="center">
  <strong>局域网文件共享工具</strong>
</div>
<br/>

<div align="center">
    <!-- platform -->
    <!-- <img alt="platform" src="https://img.shields.io/badge/platform-Linux | Windows | Mac-c3c3c3.svg?style=flat-square"/> -->
    <img alt="platform" src="https://img.shields.io/badge/-Linux-333?style=flat-square&logo=Linux&logoColor=fff"/>
    <img alt="platform" src="https://img.shields.io/badge/-Windows-0078D6?style=flat-square&logo=Windows"/>
    <img alt="platform" src="https://img.shields.io/badge/-MacOS-e67e22?style=flat-square&logo=apple&logoColor=fff"/>
</div>
<div align="center" >
	<!--version-->
    <a href="" target="blank">
		<img alt="GitHub release" src="https://img.shields.io/github/release/TCP404/OneTiny.svg?style=flat-square&color=1296DB&sort=semver"/>
    </a>
	<!--license-->
    <img alt="GitHub" src="https://img.shields.io/github/license/TCP404/OneTiny.svg?style=flat-square">
    <!--language-->
    <a href="https://golang.org" target="blank">
    	<img alt="language" src="https://img.shields.io/github/go-mod/go-version/TCP404/OneTiny/master?color=00ADD8&label=Golang&logo=Go&logoColor=00ADD8&style=flat-square"/>
    </a>
</div>
<hr />

此版本为 OneTiny 基础版，仅保留指定共享路径、端口号、是否允许上传三个选项，不带配置文件、登陆等功能。可保障最基础的文件传输功能。

若需要使用配置、登录、图形界面等，请前往 [CUI 版](https://github.com/TCP404/OneTiny) 和 [GUI 版](https://github.com/TCP404/OneTiny-gui) 进行下载。

<hr />
<br />

OneTiny 是一个用于局域网内共享文件的微型程序，它能将当前工作目录临时共享目录，对局域网内其他主机共享，通过浏览器访问 `http://局域网IP:9090` 来访问和下载共享目录中的文件。

简而言之与命令 `python -m http.server 9090` 做的是同样的事情。

## [需求]

我有两台设备，一台装着 Linux 系统，一台装着 Windows 系统，偶尔需要互相传输文件。

在 Linux 上我可以在任意一个目录下使用命令 `python -m http.server 9090`，从而在 Windows 上或局域网内其他主机上通过浏览器访问 `http://局域网IP:9090` 查看所有文件，也可以下载;

但是这条命令在 Windows 上不可行，所以需要编写一个程序可以运行在 Windows 上实现同样的功能。

## [开发技术]
- 核心功能：[gin](https://gin-gonic.com/zh-cn/)
- CLI管理：[urfave/cli](https://github.com/urfave/cli/v2)

## [使用说明]
可从本仓库的 [Release](https://github.com/TCP404/OneTiny-base/releases/) 中下载对应版本。已提供 [Linux 版](https://github.com/TCP404/OneTiny-base/releases/download/v0.3.0/OneTiny)、[Windows 版](https://github.com/TCP404/OneTiny-base/releases/download/v0.3.0/OneTiny.exe)，[Mac 版](https://github.com/TCP404/OneTiny-base/releases/download/v0.3.0/OneTiny_mac)，其他系统的同学请下载后自行编译。

### [下载]
**Linux**、**Mac**：
```bash
$ curl -LJ https://github.com/TCP404/OneTiny/releases/download/v0.3.0/OneTiny -o onetiny
$ chmod u+x ./onetiny             # 赋予执行权限
$ cp onetiny /usr/bin             # 复制到可执行文件目录并修改可执行文件名称
```

> Mac 用户注意:
> 
> 第一次打开会提示未验证开发者，可以打开 「访达」 ，打开文件所在目录，在 `onetiny` 文件处右键打开。
> 之后就可以直接用命令行运行了。

**Windows**

:point_right: [![](https://img.shields.io/github/release/TCP404/OneTiny-base.svg?style=flat-square&color=1296DB&sort=semver)](https://github.com/TCP404/OneTiny/releases/latest)

点击最新版本的 `OneTiny.exe` 进行下载

> Windows 用户注意：
> 下载时可能会有损害计算机提示，点击 「仍然保留」 即可。




### 安装（其他系统必选，Linux、Windows、MacOS 可选）
**需先安装[Golang](https://golang.org)**
```bash
$ git clone https://github.com/TCP404/OneTiny-base.git
$ go mod tidy
$ go build
```

### [Command Map]
```
onetiny
├── 不指定时可直接运行，采用默认值
├── --road   -r 指定共享目录路径
├── --port   -p 指定访问端口
└── --allow  -a 指定是否允许上传
```

### [运行]
**Windows**: 
下载后双击 `OneTiny.exe` 即可运行（需管理员权限）。
可以在CMD中切换到 `OneTiny.exe` 所在目录，执行以下任一命令：
```cmd
> OneTiny                               # 将运行在 http://本机局域网IP:9090，共享目录为当前工作目录，禁止上传
> OneTiny.exe                           # 将运行在 http://本机局域网IP:9090，共享目录为当前工作目录，禁止上传
> OneTiny -p {端口号}                    # 将运行在 http://本机局域网IP:端口号，共享目录为当前工作目录，禁止上传
> OneTiny -r {共享目录绝对路径}            # 将运行在 http://本机局域网IP:9090，共享目录为指定目录，禁止上传
> OneTiny -a [1|t|T|true|True|TRUE|空]  # 将运行在 http://本机局域网IP:9090，共享目录为当前工作目录，允许上传

> onetiny -h       # 打印帮助信息
> onetiny --help
> onetiny -v       # 打印版本信息
> onetiny --version

> onetiny -r=C:\Users\Boii -p=8192 -a -x=2      # 将运行在 「http://本机局域网IP:8192」，共享目录为 「C:\Users\Boii」 ，允许上传，允许访问共享目录往下2层
```

**Linux**、**Mac**:
```bash
$ onetiny                              # 将运行在 http://本机局域网IP:9090，共享目录为当前工作目录，禁止上传
$ onetiny -p {端口号}                   # 将运行在 http://本机局域网IP:端口号，共享目录为当前工作目录，禁止上传
$ onetiny -r {共享目录绝对路径}           # 将运行在 http://本机局域网IP:9090，共享目录为指定工作目录，禁止上传
$ onetiny -a [1|t|T|true|True|TRUE|空] # 将运行在 http://本机局域网IP:9090，共享目录为当前工作目录，允许上传

$ onetiny -h       # 打印帮助信息
$ onetiny --help
$ onetiny -v       # 打印版本信息
$ onetiny --version

$ onetiny -r=/home/boii -p=8192 -a -x=2      # 将运行在 「http://本机局域网IP:8192」，共享目录为 「/home/boii」，允许上传，允许访问共享目录往下2层
```


**更多信息**:

全局命令：
```bash
$ onetiny -h
NAME:
   OneTiny - 一个用于局域网内共享文件的FTP程序

USAGE:
   onetiny [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS] [参数...]

VERSION:
   v1.0.0

AUTHOR:
   Boii <i@tcp404.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --road  路径, -r 路径   指定对外开放的目录路径 (default: /)
   --port  端口, -p 端口   指定开放的端口 (default: 9090)
   --allow 是否, -a 是否   指定是否允许访问者上传 (default: false)

   --help,        -h        打印帮助信息 (default: false)
   --version,     -v        打印版本信息 (default: false)
```

### [访问]
1. 运行程序后，程序会提示此次服务运行在哪个端口，共享的是哪个目录，是否允许上传，允许访问的最大深度。
2. 打开浏览器，输入程序提示框中给出的地址，回车即可访问。

举个栗子：
![](README/command.png)
红色箭头所指即为服务运行地址。

打开浏览器输入程序给出的地址，即可访问共享目录中的文件：

![](README/browser.png)

**点击文件链接即可下载。**

### [结束运行]
**Linux**：
命令行中使用快捷键 `Ctrl + C` 即可停止程序运行。

**Windows**：
关闭cmd命令框即可。

> 注意：
> 1. 在 Linux 或 Mac 系统下，需要将可执行文件移动至 `/usr/bin` 下：
>     ```bash
>     mv ./OneTiny /usr/bin/tiny
>     ```
>     才能像内置命令一样使用。可以改成你喜欢的名字。
> 2. Windows 下载时可能会有损害计算机提示，点击仍然保留即可。

## [鸣谢]
- [gin](https://gin-gonic.com/zh-cn/)
- [viper](https://github.com/spf13/viper)
- [urfave/cli](https://github.com/urfave/cli/v2)
- [fatih/color](https://github.com/fatih/color)
- [schollz/progressbar](https://github.com/schollz/progressbar/v3)
- [gin-contrib/sessions](https://github.com/gin-contrib/sessions)