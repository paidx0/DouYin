# 抖音极简版

- Enter 网关入口
- user 用户相关
- feed 视频流相关
- publish 投稿发布相关
- favorite 喜欢赞相关
- comment 评论相关
- relation 粉丝关注相关

~~~text
相关依赖
    go mod tidy
安装protoc,protoc-gen-go,protoc-gen-grpc-go
    goctl env check -i -f
生成相应RPC服务
    goctl rpc protoc user.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=./rpc
生成相应API服务
    goctl api go -api user.api -dir .
~~~

~~~text
1、先要运行 
    etcd
2、启动相应的RPC服务
    go run user.go -f etc/user.yaml

~~~

## 数据库设计

- comment 评论表
    - 用户编号
    - 视频编号
    - 评论内容
- user 用户信息表
    - 用户
    - 密码
    - 关注数
    - 粉丝数
    - 点赞数
    - 用户编号
- video 视频信息表
    - 视频标题
    - 视频地址
    - 视频封面
    - 视频点赞数
    - 视频评论数
    - 文件名
    - 文件标签类
    - 视频编号
- userfocuson 用户关注表
    - 关注者编号
    - 被关注者编号
- userlikevideo 用户点赞表
    - 用户编号
    - 视频编号
- userpublishvideo 用户上传表
    - 用户编号
    - 视频编号

### 1月14日

- [x] jwt鉴权中间件
- [x] md5、uuid 小功能组件
- [x] 数据库表设计完成并定义
- [x] zap日志记录和分割完成
- [x] vip配置信息完成
- [x] MySQL、Redis配置完成
### 1月16日
- [x] 对项目结构做了较大的调整
- [x] 对数据库做了微调
- [x] 完成了日志记录的全局中间件
- [x] userRPC完成

`鉴权中间件有点问题，准备重新整合一下，其他RPC的proto做一些调整，一些可有可无的字段可以删掉`

`一个一个RPC去写，一下同时开始设计多个难免出了点错，亡羊补牢了`

### 1月18日
- [x] 完成点赞操作
- [x] 完成查看喜欢视频列表
- [x] 鉴权中间件修了
- [x] favoriteRPC完成

### 1月19日
- [x] 开始写七牛云

### 1月25日
- [x] 完成关注操作
- [x] 完成关注列表
- [x] 完成粉丝列表
- [x] 完成好友列表