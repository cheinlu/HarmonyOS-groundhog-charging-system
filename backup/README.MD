# 一、 项目介绍
go 的 web 脚手架项目。

技术栈 goframe 2.x，使用 sqlite 文件数据库。

包含功能有
- 登录，token 刷新，用户的增删改查

## 1.1 登录

登录采用 JWT 方式，登录成功生成 jwt 格式 token，其他接口会效验 token 是否有效（从 jwt 的层面，不额外查表），jwt 是分布式有效的，所以 token 不用存储，服务是无状态的。可以通过 manifest/config/config.yaml 中的配置修改 token 的存活时长。

# 二、 快速开始

## 2.1 运行

```
go run main.go
```

## 2.2 验证

本地执行 
```
curl --location 'http://127.0.0.1:8000/api/user/login' \
--header 'Content-Type: application/json' \
--data '{"username":"admin","password":"admin"}'
```

获得返回
```
{"code":0,"message":"","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg2MjE0MTgzfQ.WZV1YvwzokEgDL1urWYoUPpfWSsq3TisO3f6zutmjf4","role":""}}
```

# 二、 代码结构

## 2.1 项目框架

### 2.1.1 总体项目结构

> 项目结构分为 4个部分，api, hack, internal, manifest, 是 goframe 标准结构。

简单的介绍一下这几部分的核心功能：

- api: 定义接口的 url，输入输出字段
- hack: gf 工具配置，可以通过数据库生成通用 dao，service 代码
- internal: 业务实现代码
- manifest: 交付清单，包含程序编译、部署、运行、配置的文件。

> goframe 框架结构

| fff |  fff |
| --- | --- |
| ffff| dddd |


| 目录\文件名称 | 说明 | 描述 |
| --- | --- | --- |
| api | 对外接口 | 对外提供服务的输入/输出数据结构定义。考虑到版本管理需要，往往以api/v1...存在。 |
| hack	| 工具脚本	| 存放项目开发工具、脚本等内容。例如，CLI工具的配置，各种shell/bat脚本等文件。| 
| internal	| 内部逻辑	| 业务逻辑存放目录。通过Golang internal特性对外部隐藏可见性。| 
|   - cmd	| 入口指令	| 命令行管理目录。可以管理维护多个命令行。| 
|   - consts	| 常量定义	| 项目所有常量定义。| 
|   - controller	| 接口处理	| 接收/解析用户输入参数的入口/接口层。| 
|   - dao	| 数据访问	| 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法| 
|   - logic	| 业务封装	| 业务逻辑封装管理，特定的业务逻辑实现和封装。往往是项目中最复杂的部分。| 
|   - model	| 结构模型	| 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义。| 
|     - do	| 领域对象	| 用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改。| 
|     - entity	| 数据模型	| 数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改。| 
|   - service	| 业务接口	| 用于业务模块解耦的接口定义层。具体的接口实现在logic中进行注入。| 
| manifest	| 交付清单	| 包含程序编译、部署、运行、配置的文件。常见内容如下：| 
|   - config	| 配置管理	| 配置文件存放目录。| 
|   - docker	| 镜像文件	| Docker镜像相关依赖文件，脚本文件等等。| 
|   - deploy	| 部署文件	| 部署相关的文件。默认提供了Kubernetes集群化部署的Yaml模板，通过kustomize管理。| 
| resource	| 静态资源	| 静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中。| 
| go.mod	| 依赖管理	| 使用Go Module包管理的依赖描述文件。| 
| main.go	| 入口文件	| 程序入口文件。| 

## 2.2 业务封装

### 2.2.1 登录、鉴权

- internal/middleware/middleware.go: 提供统一鉴权处理
- internal/controller/login.go: 提供登录
- internal/controller/user.go: 提供 token 续签

# 三、 FAQ
## 1. 初始化项目失败

错误信息
```
go run main.go
# github.com/gogf/gf/contrib/drivers/sqlite/v2
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:81:12: config.Extra undefined (type *gdb.ConfigNode has no field or method Extra)
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:86:40: config.Extra undefined (type *gdb.ConfigNode has no field or method Extra)
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:150:22: undefined: gutil.GetOrDefaultStr
```

原因：gdb 版本太老，无法识别 Extra 字段。您可以尝试升级到最新版本的 gdb 包，或者尝试使用其他合适的配置来连接 SQLite 数据库，以解决此问题。

解决方案，在 go.sum 中将 github.com/gogf/gf/contrib/drivers/sqlite/v2 升级到 v2.4.0，重新 `go mod tidy`

## 2. 生成 dao 失败

错误信息
```
gf gen dao -r user           
error format "internal/dao/.go" go files: internal/dao/.go:22:3: expected 'IDENT', found '=' (and 1 more errors)
generated: internal/dao/.go
generated: internal/dao/internal/.go
error format "internal/model/do/.go" go files: internal/model/do/.go:13:7: expected 'IDENT', found 'struct' (and 1 more errors)
generated: internal/model/do/.go
error format "internal/model/entity/.go" go files: internal/model/entity/.go:12:7: expected 'IDENT', found 'struct' (and 1 more errors)
generated: internal/model/entity/.go
```

原因，命令错误，下面是正确命令，-r 是 removePrefix 的意思，把 user remove，表名就为空了

```
gf gen dao
```

## 3. go test 单元测试失败

将 _test.go 放在 controller 下，进入目录执行 go test 显示数据库驱动 sqlite 找不到。单元测试中并没有数据库连接，怀疑编译了 controller 文件。

就算 controller 增加 sqlite 驱动也不能启动。

## 4. 如何出 window 包

```
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o login.exe main.go
打包配置文件和数据库文件
zip login.zip login.exe manifest/config/config.yaml manifest/db/user.db
```


## 5. dao 根据传入的查询条件动态增加对应的查询项

通过动态增加 where 并重新赋值的方式。

``` 
model := dao.Pile.Ctx(ctx).LeftJoin("station", "station.id = pile.station_id")
if query.StationId != 0 {
    model = model.Where("pile.station_id = ?", query.StationId)
}
if query.Code != "" {
    model = model.Where("pile.code like ?", fmt.Sprintf("%%%s%%", query.Code))
}
count, err = model.Count()
```

## 6. gf 如何进行查询翻页

# 四、 附录: GoFrame 

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.