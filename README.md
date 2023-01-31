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
    goctl rpc protoc user.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
生成相应API服务
    goctl api go -api enter.api -dir .
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
