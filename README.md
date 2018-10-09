# Guardian

api测试框架，表格测试，编写json文件，即可测试api的响应以及对应数据库的正确性，以及测试并发的数据准确性。

## 下载安装

在 https://github.com/chenhg5/guardian/releases 下载对应版本，或者执行：

```bash
go get -u -v github.com/chenhg5/guardian
```

## 使用

```
guardian --tests=./example/tests/entrance.json

SUIT:  users
=================================================
获取用户                                    Ok
-------------------------------------------------
响应比对                                    ✓️
数据比对                                    ✓️
=================================================
修改用户                                    Ok
-------------------------------------------------
响应比对                                    ✓️
数据比对                                    ✓️
=================================================

Ok

SUIT:  orders
=================================================
获取订单                                    Ok
-------------------------------------------------
响应比对                                    ✓️
数据比对                                    ✓️
=================================================
创建订单                                    Ok
-------------------------------------------------
响应比对                                    ✓️
数据比对                                    ✓️
=================================================

Ok

```

```entrance.json```为入口文件，例子详见 example/tests/entrance.json

## json格式

### 入口文件

| 选项名 | 子选项名 | 描述 | 格式 | 是否必须 | 例子 |
| ------ | ------ | ------ | ------ | ------ | ------ |
| database | port | 端口 |  字符串 | 是 | 3306 |
|  | user | 用户 | 字符串 | 是 | root |
|  | password | 密码 | 字符串 | 是 | root |
|  | charset | 字符集 | 字符串 | 是 | utf8 |
|  | host | 地址 | 字符串 | 是 | 127.0.0.1 |
|  | database | 数据库名 | 字符串 | 是 | guardian |
| tables | 无 | 案例集 | 对象 | 是 | { "users": [ "users/get.json", "users/post.json" ], "orders": ["orders/get.json","orders/post.json"]
| vars | 无 | 全局变量 | 对象 | 否 | { "host": "http://127.0.0.1:1235" } |
                            
### 测试案例

| 选项名 | 子选项名 | 描述 | 格式 | 是否必须 | 例子 |
| ------ | ------ | ------ | ------ | ------ | ------ |
| info | title | 标题 | 字符串 | 是 | 这是一个标题 |
|  | description | 描述 | 字符串 | 是 | 这是一个描述 |
| concurrent | 无 | 并发数 | 整数 | 否 | 1 |
| request | url | url | 字符串 | 是 | {{host}}/user |
|  | method | 方法，仅提供四种: json, get, formget, formpost | 字符串 | 是 | get |
|  | params | 参数 | 对象或数组 | 否 | {"id": 1} |
|  | header | 头部 | 对象  | 否 | {"token": "1231313"} |
| response | body | 返回body数据 | 字符串或对象 | 是 | 123 |
|  | header | 头部 | 对象  | 否 | {"token": "1231313"} |
| data | 无 | 验证数据 | 数组 | 否 | [<br>{<br>"sql": "select name from user where id = 1",<br>"result": [<br>{<br>"name": "jack"<br>}<br>]<br>}<br>] |

## 项目信息

### 流程

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
    
### 组成部分

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

### 表格读取

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