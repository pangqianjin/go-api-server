# go-apiServer

#### 介绍
GO语言编写的静态资源服务器，也可用于mock JSON请求，自用

#### 软件架构
使用全宇宙最快的go框架Iris编写


#### 安装教程

1.  $```go get github.com/kataras/iris/v12@master # or @v12.2.0-beta7```

#### 使用说明

1. property-config.json请参考默认的配置
```json
{
    "port": "9999",
    "data_path": "./mockData",
    "public_path": "/Users/pangqianjin/Public/fqfin/server-public/resource/statics",
    "debug_level": "info"
}
```
2. 默认使用9999端口，如冲突请修改为其他即可
3. 默认配置./mockData为JSON文件夹, 最多支持三层嵌套/xx/xx/xx[.post].json
    - 拷贝你的JSON文件夹到mockData下，如 /mockData/api/*.json
    - \*.json可用于GET,PUT,DELETE请求，*.post.json用于POST请求
4. 静态资源目录请配置public_path字段
5. 日志等级请配置debug_level字段，默认为info
6. 日志自定义请配置logger-config.json
6. 本人上传的./go-apiServer为macos m1可用，其他操作系统请自行编译

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
