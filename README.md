# golang学习笔记（快速入门）

个人Go语言学习demo 本项目demo 均基于

- [8小时 语雀笔记](https://www.yuque.com/aceld/mo95lb/ovib08)
- [8小时 B站视频](https://www.bilibili.com/video/BV1gf4y1r79E)
- [zinx 框架项目地址](https://gitee.com/Aceld/zinx)

以下内容均来自与对 **刘丹冰Aceld** 大佬的视频和笔记的整理

感谢 **刘丹冰Aceld** 大佬的无私奉献！

# 一、学习要求

- 非0基础 需要有其他语言开发的基础(比如 C/Java/php 等偏后端语言环境)
- 基本网络编程(TPC/IP socket 等) 和 并发思想 
- 计算机基本认知(系统 服务器 硬件 等)
- 需要有Linux基础(常用指令)

## 适合人群

已经掌握一门语言 比如Java后端 需要拓展知识栈 快速上手golang 的人群

## 大纲
- golang环境安装
- golang语言特性
- golang语法
- golang进阶
- golang实战案例

# 二、golang 开发环境安装

刘大佬是Linux+vscode 需要Linux开发环境的可以看刘大佬的视频

视频：[开发环境重点和IDE选型推荐 P2](https://www.bilibili.com/video/BV1gf4y1r79E/?p=2)

因为我日常开发环境都是windows只有部署项目时会部署的linux，所以我的环境是windows。

VSCode和GoLand都有用，不过对我来说我只是吧VSCode当文本编辑器来用，并不熟悉，日常开发AndroidStudio和IDE用的比较多，GoLand操作起来更顺手，所以选用GoLand。

Goland破解请自行百度，有能力的富哥可以支持正版。

## Go下载

[golang 镜像站 下载地址](https://golang.google.cn/dl/)

![golang_download.png](img/golang_download.png)

下载完成后 疯狂 **下一步** 即可

## 配置环境

打开环境变量 配置GoPath 和 GoFile

## 验证

打开cmd 输入查看安装的版本

```shell
go version
```

如果出现提示

```shell
go version go1.20.4 windows/amd64
```

可以看到我这里安装的是 1.20.4的版本 说明安装并且配置没问题

然后安装GoLand

## GoLand下载

[GoLand官网下载地址](https://www.jetbrains.com/go/)

![goland_download.png](img/goland_download.png)

# 三、golang语言特性

## 1、Golang的优势

### 1) 极简的部署方式
- 可直接编译成机器码
- 不依赖其他库
- 直接运行即可部署

编译演示

进入某个项目的目录
```shell
go build main.go
```
编译成功会生成一个 main 可执行文件

查看这个main用了哪些依赖库
```shell
ldd main
```
可以看到它只包含标准的 so库 libc库 和 pthread库 不依赖其他任何库

可使用 ./ 命令 直接执行
```shell
./ main
```

以上演示为Linux环境

windows环境下编译同理

```shell
go build main.go
```

在windows下编译成功后会生成 **main.exe** 可执行程序

我们双击运行即可 或者 在cmd里运行也行 

**一般我们开发调试在IDE里操作即可 不用专门使用命令行编译 这里只是演示过程**

### 2) 静态类型语言

静态语言(Java, C, Go 等)的优势：可以在编译时检查出隐藏的问题(比如哪一行语法错误 或者少写了个符号等) 和静态语言对立的是动态语言(如shell脚本, python脚本, js等)，动态语言没有编译器 所以只能在程序运行过程中逐条判断

### 3) 语言层面的并发

- 天生支持并发 非包装类型的优化
- 充分的利用多核 尽量提高CPU的利用率

### 4) 强大的标准库

- runtime 系统调度机制
- 高效的GC垃圾回收
- 丰富的标准库 (文本操作 输入/输出 时间/日期 json/xml socket/rpc 线程/锁 文件系统 并发 邮件 加解密 等覆盖了日常开发绝大部分场景)

### 5） 简单易学

- 仅有 25 个关键字
- C语言语法简洁 内嵌C语法支持
- 面向对象特征(继承 封装 多态)
- 跨平台语言 类似于java 只要装了go环境 都可以运行

### 6) 大厂领军

国内外各种大厂都在用 也有很多开源框架 给你用

## 2、Golang适合做什么(强项)

### 1) 云计算基础设施领域

docker, kubernetes, etcd, consul, cloudflare CDN, 七牛云 等

### 2) 基础后端软件

tidb, influxdb, cockroachdb 等

### 3) 微服务

go-kit, micro, monzo/typhon, bilibili 等

### 4) 互联网基础设施

以太坊, hyperledger 等

## 3、Golang大作

- Docker (16年docker火了 让go成了16年 年度语言)
- kubernetes

## 4、Golang的不足

### 1) 包管理

我们所使用大部分第三方库 都是托管在github上的

### 2) 无泛型 (目前已经支持)

早期go不支持泛型 不过在1.18已经支持了 

### 3) Error

所有的 **Exception** 都用 **Error** 来处理(比较有争议)

没有java的 try-catch 捕获异常的操作 只能 error 一层一层抛出

### 4) 对C的降级处理

并非无缝，没有C降级到asm那么完美(序列化问题)

兼容c 不是完全兼容 可以导入 c的包 调用c的函数

# 语法

## 1、hello world

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

执行

```shell
go run main.go
```

## 2、变量

### 1) 先声明 再赋值

```go
//先声明
var str string
var i int
//再赋值
str = "ddd"
i = 188
fmt.Print("我是 str: ", str,"我是 i: ",i)
```

### 2) 自动判断类型
```go
var b = false
var i2 = 99
fmt.Print("我是 b: ", b, "我是 i2: ", i2)
```

### 3) 省略var 直接 :=

```go
dev := "Go Developers"
fmt.Println("Go Hello World! ", dev)
```

## 2.5、 多变量声明

```go

```




## 3、常量



