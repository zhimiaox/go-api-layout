# 纸喵golang api后端服务专用模板

#### 介绍

没啥可介绍的, member为完整示例，但因为需要调用数据库，因此并没有挂载，如若开启，请到api/api.go中开启，并完善config.toml配置文件

#### 开发文档

> Api文档自动生成

```shell
go get -u github.com/swaggo/swag/cmd/swag

swag init --parseDependency --parseVendor --parseInternal
```

> 项目运行

1、配置golang环境

2、打开项目，直接运行`man.go`

3、 swag地址 http://localhost:1324/docs/index.html

4、 测试地址 http://localhost:1324/api/v1/hello

> 开发规约

1、develop分支挂载ci/cd，每次提交都会自动编译部署，因此要求每次提交必须可以编译运行，若遇到临时提交，请单独创建临时分支

> 任务管理