# 简介
 一款使用 Gin + Gorm 搭建的基础 Go API 框架 

# 目录结构

```
.
├── app                                                             // 应用代码
│   ├── controllers                                                 // 控制器
│   ├── middlewares                                                 // 中间件
│   ├── models                                                      // 模型
│   ├── requests                                                    // 请求验证
│   ├── services                                                    // 服务
├── bootstrap                                                       // 框架引导
├── cmd                                                             // 命令
├── config                                                          // 配置
├── database                                                        // 数据库
│   ├── factories                                                   // 数据工厂
│   ├── migrations                                                  // 数据迁移
│   ├── seeders                                                     // 数据填充
├── pkg                                                             // 公共包
├── resources                                                       // 资源
├── routes                                                          // 路由
├── storage                                                         // 本地存储
├── .env                                                            // 环境变量
├── main.go                                                         // 入口文件
├── go.mod
├── go.sum
├── README.md
```

# 初始化

执行命令生成环境变量文件，根据项目修改 `.env` 文件配置参数

```shell
cp .env.example .env
```

执行命令生成项目 `app_key`， 用作加密，需把值手动添加进 `.env` 文件 `APP_KEY` 环境变量

```shell
go run main.go key
```

# 命令行工具

框架自带一些命令，可以节约一些开发时间，如下：

生成控制器文件
```shell
go run main.go make controller user
```

生成模型文件
```shell
go run main.go make model user
```

生成请求验证文件
```shell
go run main.go make request user
```

生成命令行文件
```shell
go run main.go make cmd sync_user
```

生成数据迁移文件
```shell
go run main.go make migration add_users_table user
```

执行未迁移过的文件
```shell
go run main.go migrate up
```

回滚上一次操作的迁移
```shell
go run main.go migrate rollback
```

回滚所有迁移
```shell
go run main.go migrate reset
```

回滚所有迁移，并执行全部迁移
```shell
go run main.go migrate refresh
```

删除所有表，并执行全部迁移
```shell
go run main.go migrate fresh
```