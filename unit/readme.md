
# Go测试框架 testify 10k star
- https://github.com/stretchr/testify

|子包|功能|测试案例|
|mock	|模拟实现|模拟依赖服务|
|assert |断言测试案例|ok|
|http	|http 测试||
|require|断言||

# interface mock 生成器生成模仿对象
- https://github.com/vektra/mockery
- 使用方式 mockery --all --keeptree

# Go测试框架goconvey 6k star
- github.com/smartystreets/goconvey/convey

1. convey 测试单元
2. so 	  断言

# 数据库mock sql mock 2.6k
- https://github.com/DATA-DOG/go-sqlmock

# 最终目的

|package|解决方法|备注|
|----------|------|-------
|repository|数据库 sqlmock||
|usecase|testify mock|模拟依赖，进行测试|
|deliver|testify http|


# 测试驱动开发(Test-Driven Development, TDD)

# 测试代码原则
- A：Automatic（自动化）
- I：Independent（独立性）
- R：Repeatable（可重复）



参考文章：
- sqlmock https://blog.marvel6.cn/2020/01/test-and-mock-db-by-xorm-with-the-help-of-convey-and-sqlmock/

