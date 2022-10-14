# jkframe

[![jkframe](https://github.com/jkstack/jkframe/actions/workflows/test.yml/badge.svg)](https://github.com/jkstack/jkframe/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/jkstack/jkframe.svg)](https://pkg.go.dev/github.com/jkstack/jkframe)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkstack/jkframe)](https://goreportcard.com/report/github.com/jkstack/jkframe)
[![go-mod](https://img.shields.io/github/go-mod/go-version/jkstack/jkframe)](https://github.com/jkstack/jkframe)
[![license](https://img.shields.io/github/license/jkstack/jkframe)](https://opensource.org/licenses/MIT)

[English](README.md) | [中文文档](README_zh.md)

精鲲golang项目基础组件封装库

## 模块分类

* [api](/api): http服务端常用接口封装，包含参数处理和返回内容处理等接口
* [l2cache](/cache/l2cache): 通过内存和磁盘文件实现的二级缓存，当写入内容超过预定值时缓存内容将被刷写到磁盘上
* [kvconf](/conf/kvconf): 对于常见的kv形式配置文件进行处理，支持kv形式配置文件的读取和生成
* [yaml](/conf/yaml): 扩展yaml基础库，支持#include语法可将一个复杂的yaml文件拆分成多个文件
* [daemon](/daemon): 守护进程封装，支持启动子进程时切换用户身份，子进程异常退出自动重启等
* [logging](/logging): log模块，支持log分级并输出不同格式的log，支持按照日期或日志文件大小自动切分等
* [mysqlschema](/mysqlschema): mysql数据库表结构和初始化数据生成库，支持按照版本号进行表结构升级以及注入初始化数据
* [utils](/utils): 一些常用方法，如生成随机字符串、assert、获取当前调用堆栈信息等
* [stat](/stat): 基于prometheus的微服务埋点库
* [compress](/compress): 通用的数据压缩算法

## 使用方法

    go get -u github.com/jkstack/jkframe

## 文档

GoDoc: https://pkg.go.dev/github.com/jkstack/jkframe

## TODO

1. prometheus埋点库封装
2. 支持restful风格的API接口