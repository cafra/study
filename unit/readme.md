# 单元测试哪些问题

## 组件依赖
- 组件依赖，通常是数据库，需要链接环境，造数据，才可以完成测试，环境成本高
- 下次测试时，需要重新模拟数据，完成测试

## 服务依赖
- 开发服务A时，依赖服务B，此时测试A时，依赖服务B联调测试，代码覆盖度很难保证，在微服务中更明显；
- 模拟依赖的不同case也很麻烦

## 测试单元迭代困难
- 需求迭代或者代码重构后，无法快速自动测试

# 测试代码原则
- A：Automatic（自动化）
- I：Independent（独立性）
- R：Repeatable（可重复）

# 目前的单元测试方式，基本不符合这些原则

## 如何解决这些问题 思想: mock

## 具体数据实践方案

|package|解决方法|备注|状态|
|----------|------|-------
|repository|数据库 sqlmock|模拟数据库|Y|
|usecase|testify mock|模拟依赖，进行测试|Y|
|deliver|testify http|模拟依赖，进行测试|Y|

## Go测试框架 testify 10k star

- https://github.com/stretchr/testify

|子包|功能|测试案例|
|mock	|模拟实现|模拟依赖服务 mock case|
|assert |断言测试案例|ok|
|http	|http 测试||
|require|断言||

## interface mock 生成器生成接口对象

- https://github.com/vektra/mockery
- 使用方式 mockery --all --keeptree

## Go测试框架goconvey 6k star

- github.com/smartystreets/goconvey/convey

|packge|
|-------|------|
|convey| 测试单元|
|so 	  |断言|

## 数据库 mock 2.6k

- https://github.com/DATA-DOG/go-sqlmock

## Struct Data Fake Generator

- https://github.com/bxcodec/faker


参考文章：
- sqlmock https://blog.marvel6.cn/2020/01/test-and-mock-db-by-xorm-with-the-help-of-convey-and-sqlmock/

