# jkframe

[![jkframe](https://github.com/jkstack/jkframe/actions/workflows/test.yml/badge.svg)](https://github.com/jkstack/jkframe/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/jkstack/jkframe.svg)](https://pkg.go.dev/github.com/jkstack/jkframe)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkstack/jkframe)](https://goreportcard.com/report/github.com/jkstack/jkframe)
[![go-mod](https://img.shields.io/github/go-mod/go-version/jkstack/jkframe)](https://github.com/jkstack/jkframe)
[![license](https://img.shields.io/github/license/jkstack/jkframe)](https://opensource.org/licenses/MIT)

[English](README.md) | [中文文档](README_zh.md)

jkstack golang basic framework

## modules

* [api](/api): http handler interface, includes arguments parser and response reply
* [l2cache](/cache/l2cache): cache with memory and disk, when write data size more than values seted, the cache data will write to the disk
* [kvconf](/conf/kvconf): supported configure file with key=value
* [yaml](/conf/yaml): supported yaml file parse, added #include keyword to include another yaml file
* [daemon](/daemon): daemon process supported, supported change run user in child process, supported restart when child process exited
* [logging](/logging): log library, supported log level of each format, supported log rotate by date and log file size
* [mysqlschema](/mysqlschema): mysql table schema and initialize data management tools, supported table schema version and initialize data version
* [utils](/utils): common function, like UUID, assert, get callstack
* prometheus: prometheus library(developing)

## usage

    go get -u github.com/jkstack/jkframe

## documentation

GoDoc: https://pkg.go.dev/github.com/jkstack/jkframe