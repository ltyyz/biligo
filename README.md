## BiliGo

BiliGo 是一个由多个 Go 开源库组装而成的快速开发 webapi 的系统，
目前还处于初级阶段，功能性和稳定性没有任何保障

## 为什么有 BiliGo？

1. 强大的框架太复杂，没必要
2. 简单的框架太简单，不够用
3. BiliGo 的目标是上手即用，像呼吸

## 使用到的库

- [Gin](https://github.com/gin-gonic/gin) 轻量路由框架
- [logrus](https://github.com/sirupsen/logrus) 日志框架
- [goprofile](https://github.com/ltyyz/goprofile) 多环境配置文件管理工具
- [gorm](https://github.com/jinzhu/gorm) 数据库 orm 框架
- [uuid](github.com/gofrs/uuid) uuid 生成库

## 文件夹介绍及主要文件

```
├─main.go               # 程序 main 入口
├─config                # 配置文件 goprofile
├─constant              # 用到的常量
├─mysql                 # MySQL 连接
├─log                   # 日志 logrus
├─util                  # 工具代码
└─modules               # 接口代码
    ├─app               # **你的代码应该在这里**
    │  ├─api            #   uri 入口
    │  ├─model          #   数据库 ORM 对象文件
    │  ├─service        #   service 代码
    │  └─router.go      #   模块路由设置
    ├─system            # BiliGo 内置系统模块（可能会写些东西）
    │   ├─api
    │   ├─model
    │   ├─service
    │   └─router.go
    ├─auth              # 登录认证模块
    └─router.go         # 总路由

```

## 快速上手

`go version 1.12.9`

创建数据库，并导入脚本 `mysql/init.sql`

修改配置文件 `config/config.env` 中的数据库相关配置 

(环境变量已有可以不用指定)

```
export GOPROXY=http://mirrors.aliyun.com/goproxy/
export GO111MODULE=on

go run .
```

然后浏览器访问：

```
http://localhost:8888/
```

就可以看到结果：

```json
{
    "code": 0,
    "message": "处理成功",
    "success": true,
    "data": "Hello BiliGo",
    "timestamp": 1566017186754
}
```

## 最后

BiliGo 的目标是编写 api ，所以不会加入模板处理
