# NanoOJ

尝试一个OJ系统的手写

## 项目概览

目前是纯C/C++ OJ。整个项目使用了 Vue全家桶 + ElementPlus + Golang-Gin + C++ 开发。

>（其实本来想直接用docker当判题机环境的，但是还没做）

- `judge` 判题端，使用 GNU/Linux syscalls 和 C++ 完成，目前还没加输入支持和防止禁用源码系统调用的功能，安全性尚不能用于生产环境。
- `web` 网页端。使用 Vue+ElementUIPlus 搓的单页应用(SPA)。目前完成了大半~~才怪~~，剩下一些~~一堆~~细节。

>想提前试下判题机可以直接把 `judge/core/` 里边那些玩意编译出来（要用Linux编译哦）用用看。
>用法的话，不带参数直接运行程序就会输出。

```tree
NanoOJ/
├── judge/  # 后端目录
│   ├── config/  # 配置文件目录
│   │   ├── config.go  # 读取和处理配置文件的代码
│   │   └── config.yaml  # 配置文件，例如数据库连接信息
│   ├── controller/  # 控制器目录
│   │   ├── auth.go  # 处理用户身份验证相关请求的代码
│   │   ├── problem.go  # 处理问题相关请求的代码
│   │   ├── submission.go  # 处理代码提交相关请求的代码
│   │   ├── testset.go  # 处理测试集相关请求的代码
│   │   └── user.go  # 处理用户相关请求的代码
│   ├── database/  # 数据库目录
│   │   ├── migrate/  # 数据库迁移文件目录
│   │   │   ├── 20220301_init.sql  # 初始化数据库表的 SQL 文件
│   │   │   └── 20220324_add_submission_status_column.sql  # 增加提交状态列的 SQL 文件
│   │   ├── model/  # 数据模型目录
│   │   │   ├── problem.go  # 问题模型
│   │   │   ├── submission.go  # 提交模型
│   │   │   └── user.go  # 用户模型
│   │   ├── database.go  # 数据库连接和初始化代码
│   │   └── repository.go  # 数据库操作的统一接口
│   ├── middleware/  # 中间件目录
│   │   ├── auth.go  # 身份验证中间件
│   │   └── logging.go  # 日志中间件
│   ├── router/  # 路由目录
│   │   ├── middleware.go  # 路由中间件
│   │   └── router.go  # 路由处理代码
│   ├── worker/  # 后台工作目录
│   │   └── judger.go  # 评测机代码
│   ├── main.go  # 后端程序入口文件
│   └── go.mod  # Golang 依赖管理文件
│
└── web/  # 前端目录
    ├── public/  # 静态资源目录
    ├── src/  # 代码目录
    │   ├── api/  # API 相关代码目录
    │   │   ├── auth.js  # 处理用户身份验证相关 API 的代码
    │   │   ├── problem.js  # 处理问题相关 API 的代码
    │   │   ├── submission.js  # 处理提交相关 API 的代码
    │   │   ├── testset.js  # 处理测试集相关 API 的代码
    │   │   └── user.js # 
    │   ├── assets/  # 静态资源目录
    │   ├── components/  # 组件目录
    │   ├── router/  # 路由目录
    │   ├── store/  # Vuex 相关代码目录
    │   │   ├── modules/  # Vuex 模块目录
    │   │   ├── actions.js  # Vuex actions 代码
    │   │   ├── getters.js  # Vuex getters 代码
    │   │   ├── index.js  # Vuex 入口文件
    │   │   ├── mutations.js  # Vuex mutations 代码
    │   │   └── state.js  # Vuex state 代码
    │   ├── views/  # 视图目录
    │   ├── App.vue  # Vue 根组件
    │   ├── main.js  # Vue 程序入口文件
    │   ├── router.js  # Vue 路由配置文件
    │   └── store.js  # Vuex 状态管理文件
    ├── .eslintrc.js  # ESLint 配置文件
    ├── babel.config.js  # Babel 配置文件
    ├── package.json  # NPM 依赖管理文件
    ├── postcss.config.js  # PostCSS 配置文件
    ├── README.md  # 项目说明文件
    ├── vue.config.js  # Vue CLI 配置文件
    └── yarn.lock  # Yarn 依赖管理锁文件
```

## 计划

- 先把数据库设计好吧（摊
- 啊，然后是前后端功能对接，也稍微再完善一下吧
- 然后嘛......还没有写管理员相关的功能，也稍微写一下吧
- 话说 `web/src/views/` 和 `components/` 乱的一团糟，那也稍微收拾下吧
- 还有好多好看的 ElementUIPlus 组件没用呢
- 话说判题端也得加上读取stdin的功能了
- 以及得想办法禁了源码的系统调用，不然服务器一定会炸的
- 多线程支持也加上吧，大概能快一些吧
- ~~总感觉好累啊 那就有时间再摸吧~~
