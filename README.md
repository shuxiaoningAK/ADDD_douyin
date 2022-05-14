# ADDD_DOUYIN
This is a project for a group. The goal is for the champion.
这是一个团队工程。

**实现的一个mini版抖音服务端**

## 1.基础接口
   1. - [x]  视频流接口   //TODO开发中
   2. - [x]  用户注册
   3. - [x]  用户登录
   4. - [ ]  用户信息
   5. - [ ]  投稿接口
   6. - [ ]  发布接口

## 2.扩展接口 - I
   1. - [ ]  赞操作
   1. - [ ]  点赞列表
   2. - [ ]  评论操作
   3. - [ ]  评论列表
   
## 2.扩展接口 - II
   1. - [ ]  关注操作
   1. - [ ]  关注列表
   2. - [ ]  粉丝列表




## 项目结构

```shell
douyin/
├── conf
├── controller
├── middleware
├── model
├── routes
├── serializer
├── service
└── util
```
![](./summ.png)

- conf : 用于存储配置文件
- controller : 用于处理请求JSON格式返回
- middleware : 自定义的中间件
- model : gorm使用的数据库模型
- routes : 路由转发
- serializer：定义序列化json的结构体
- service : 具体的功能处理逻辑


## 如何运行本项目

**1.下载第三方库**
```go
go mod tidy
```
**2.配置数据库**
项目已经使用gorm开启了**数据库自动迁移模式**，仅需在conf/config.ini文件中配置好参数，开启数据库即可。

**3.运行**
```go
go run main.go
```
运行项目生成数据表后，请运行以下SQL语句使得Gorm大小写敏感
```SQL
ALTER TABLE user MODIFY COLUMN name VARCHAR(50) BINARY CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL;
```
**4.测试**
1. 可以使用**postman**对接口进行测试
![](postman.png)
2. 也可使用客户端测试


## 后续的优化
TODO
是否考虑过SQL注入，越权等安全问题的防御方式
数据表是否设置了合理的索引，处理了常见的性能问题
等等。。。

