# NanoOJ

使用`vue3`+`ts`+`element-plus`+`go`+`gorm`+`gin`实现的OJ系统

## Doc

整个项目使用了 Vue全家桶 + ElementPlus + Golang-Gin + C++ 开发。

- `judge` Web 服务器+判题端实现，使用 Go + C++ 完成。
- `web` 网页端。使用 Vue3+ElementPlus 搓的单页应用。

```bash
NanoOJ/
├── judge
│   ├── config
│   │   ├── config.go
│   ├── core
│   │   ├── Dockerfile
│   │   └── judge.cc
│   ├── database
│   │   └── repository.go
│   ├── model
│   ├── worker
│   ├── lib
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── LICENSE
├── Makefile
├── README.md
└── web
    ├── src
    │   ├── api
    │   ├── assets
    │   ├── components
    │   ├── utils
    │   ├── App.vue
    │   ├── main.ts
    │   ├── model.ts
    │   └── views
    └── vite.config.ts
```

开发顺序：功能梳理->数据库设计->后端设计+前端设计->编写->测试->部署。功能设计好后，以TDD的方式进行数据库设计和后端开发。

功能主要是几个数据模型的交互，包括用户，题目，提交记录，练习/竞赛，排行榜等。涉及数据库的设计和数据表关系的维护。

系统架构是典型前后端分离，网页作为客户端通过API和服务器交互，服务器负责数据库层操作，以及判题机的调度。

### Judger

后端可以是单体/分布式架构，一个实例可以作为**前端-数据库**处理，也可以负责**判题机-数据库**处理。

系统实现了多种数据库接口兼容，默认使用内置的SQLite3，也支持MySQL。

后端的判题系统包含一个判题镜像，这镜像中包含了判题核心程序，以及编译环境。

判题机的设计：一个可执行文件，它的输入是一个目录，这个目录中包含了测试用例和用户代码。**判题机按照规定格式输出判题结果**到stdout，供后端读取。或者也可以直接输出到一个文件中，供后端读取。

允许用户上传**实现特定API**的判题机，创建判题机镜像。创建实例的时候，挂载待判定的目录，以及测试用例的目录，来让后端直接判定固定目录中的内容，实现镜像自动判题。

镜像实例全部保持断网，限制判题机实例的资源占用情况。

>尝试支持OS-Lab

### Web

典型MVC单页应用。用函数式的方法提高了抽象度，减少了代码量。

#### 主页

没啥说的，就是几个组件拿出来拼一下，整个状态展示和公告栏。

#### 题库

所有题目列表。顺道，一些题目可以从主题库隐藏，因此可以给题目加上所属组的属性：属于默认组，就会出现在这里。同时，也可以给题目增加标签设计，便于分类查看。另外也可以加上搜素功能。

#### 排行榜

也是个可复用组件，有总榜（根据用户积分决定），对于每个练习/竞赛有分榜，对于每个题目也有分榜。另外加上封榜功能，为OI赛制提供支持。

#### 题目详情

时间/空间限制，提交代码/提交文件，支持填空题。展示题目的历史提交记录展示列表，以及当前状态（分数，是否通过等），其他信息（问题描述）的展示。

#### 练习/竞赛详情

一个练习有题目列表，练习信息、排行榜等几个组件。封榜，做题进度，提交记录查看。

## TODO

- [x]i18n
- [x]无Docker环境下的判题机&部署文档
- [x]UI细节修复
- [x]响应式布局
- [ ]负载均衡
- [x]数据导入/导出
- [x]token请求时无感刷新
- [x]封榜
- [x]修复各种非标准POST请求：submission, contest等

## Changelog

- 2022.3: 项目启动
- 2023.3: 项目demo完成
- 2024.3: 详细实现&代码优化
    2024.3.15: 重构完成，实现Docker判题机外的大多数功能
    2024.3.17: 权限系统重构完成
    2024.3.19: 本地C/C++判题机测试完成

## LICENSE

GNU General Public License V3
