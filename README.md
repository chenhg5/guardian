# Guardian

api测试框架，表格测试，编写json文件，即可测试api的响应以及对应数据库的正确性，以及测试并发的数据准确性。

## 使用

```
go get -u -v github.com/chenhg5/guardian
go install github.com/chenhg5/guardian

guardian --tests=./example/entrance.json

SUIT: order
=================================================
create_order                        OK
=================================================
pay_order                          Fail
=================================================

Fail

```

```entrance.json```为入口文件，例子详见 example/entrance.json

## 流程

- 输入表格集
    - 输入参数：
        - url
        - method
        - param
    
    - 对比响应：
        - json
        - 字符串
    
    - 核对数据：
        - mysql
        - redis
    
    - 记录结果        
- 输出结果
    
## 组成部分

- 请求函数
    - json post
    - get
    - 表单 post
    - 表单 get
    
- 响应对比
    - 字符串对比
        - 完全一致
        - 包含
    - json对比
        - 完全一致
        - 结构一致
        - 结构一致且部分数据一致

- 数据获取
    - sql连接
        - mysql
    - redis

- 全局变量管理
    - 定义全局变量
    - 读取变量  

- io部分
    - 输出到终端/文件
    - 读取文件

## 表格读取

入口文件为json格式

- 数据库
- redis
- 测试表格集

表格为json格式

- 测试信息
    - 标题
    - 描述

- 请求
    - url
    - method
    - param
    - header   
    
- 响应
    - header
    - body
    - 对比方式

- 数据
    - sql
    - 期待结果
    - 对比方式                               