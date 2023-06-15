# tkl

#### 项目概述

二手交易网站

#### 接口文档

https://apifox.com/apidoc/shared-344f5ea8-7a45-4270-99db-1b293b9c0053

#### 安装指南

1. 安装步骤：分别移步backend和frontend目录查看README.md

2. 环境要求：
   
   - Go 1.20
   
   - Node.js 18.16.0（16.0.0以上）

#### 运行指南

分别移步backend和frontend目录查看README.md

#### 贡献指南

1. clone 本仓库
2. commit 代码
3. push 代码自动生成 pull request
4. 管理员审查代码并合并

#### 技术栈

1. 后端  
   go + gin + gorm + mysql
2. 前端  
   vue3 + vite +  pinia + vue router  + element plus + axios

#### 代码目录结构

```
|-- tkl
 |-- LICENSE
 |-- README.md
 |-- 简单需求分析.md
 |-- backend
 | |-- go.mod
 | |-- go.sum
 | |-- main.go
 | |-- README.md
 | |-- common
 | | |-- database.go
 | | |-- jwt.go
 | |-- controller
 | | |-- BannerController.go
 | | |-- CartController.go
 | | |-- ChatController.go
 | | |-- GoodsController.go
 | | |-- ImageController.go
 | | |-- OrderController.go
 | | |-- UserController.go
 | |-- middleware
 | | |-- AuthMiddleware.go
 | | |-- CORSMiddleware.go
 | | |-- RecoveryMiddleware.go
 | |-- model
 | | |-- banner.go
 | | |-- Cart.go
 | | |-- category.go
 | | |-- chat.go
 | | |-- chatList.go
 | | |-- commodity.go
 | | |-- goods.go
 | | |-- image.go
 | | |-- order.go
 | | |-- picture.go
 | | |-- user.go
 | | |-- userAddress.go
 | |-- routes
 | |-- routes.go
 |-- frontend
 |-- index.html
 |-- jsconfig.json
 |-- package-lock.json
 |-- package.json
 |-- README.md
 |-- vite.config.js
 |-- public
 | |-- favicon.ico
 |-- src
 |-- App.vue
 |-- main.js
 |-- apis
 | |-- cart.js
 | |-- category.js
 | |-- chat.js
 | |-- checkout.js
 | |-- detail.js
 | |-- home.js
 | |-- image.js
 | |-- layout.js
 | |-- likelist.js
 | |-- login.js
 | |-- order.js
 | |-- pay.js
 | |-- register.js
 | |-- release.js
 | |-- testAPI.js
 | |-- userInfo.js
 |-- assets
 | |-- base.css
 | |-- logo.svg
 | |-- main.css
 | |-- images
 | |-- img
 | |-- emoji
 |-- components
 | |-- HelloWorld.vue
 | |-- TheWelcome.vue
 | |-- WelcomeItem.vue
 | |-- icons
 | | |-- IconCommunity.vue
 | | |-- IconDocumentation.vue
 | | |-- IconEcosystem.vue
 | | |-- IconSupport.vue
 | | |-- IconTooling.vue
 | |-- imageView
 | |-- index.vue
 |-- composables
 | |-- useCountDown.js
 |-- directives
 | |-- index.js
 |-- router
 | |-- index.js
 |-- stores
 | |-- cartStore.js
 | |-- category.js
 | |-- counter.js
 | |-- user.js
 |-- styles
 | |-- common.scss
 | |-- var.scss
 | |-- element
 | |-- index.scss
 |-- utils
 | |-- http.js
 |-- views
 |-- CartList
 | |-- index.vue
 |-- Category
 | |-- index.vue
 | |-- composables
 | |-- useBanner.js
 | |-- useCategory.js
 |-- Chat
 | |-- index.vue
 | |-- components
 | |-- chathome.vue
 | |-- chatwindow.vue
 | |-- HeadPortrait.vue
 | |-- Nav.vue
 | |-- PersonCard.vue
 |-- Checkout
 | |-- index.vue
 |-- Detail
 | |-- index.vue
 |-- Home
 | |-- index.vue
 | |-- components
 | |-- GoodsItem.vue
 | |-- HomeBanner.vue
 | |-- HomeCategory.vue
 | |-- HomeHot.vue
 | |-- HomeNew.vue
 | |-- HomePanel.vue
 | |-- HomeProduct.vue
 |-- Layout
 | |-- index.vue
 | |-- components
 | |-- HeaderCart.vue
 | |-- LayoutFixed.vue
 | |-- LayoutFooter.vue
 | |-- LayoutHeader.vue
 | |-- LayoutNav.vue
 |-- Login
 | |-- index.vue
 |-- Member
 | |-- index.vue
 | |-- components
 | |-- SoldOrder.vue
 | |-- UserInfo.vue
 | |-- UserOrder.vue
 |-- Pay
 | |-- index.vue
 | |-- PayBack.vue
 |-- Register
 | |-- index.vue
 |-- Release
 |-- index.vue
```

#### 示例和演示

#### 软件架构

后端架构

1. 编程语言：Go

2. 框架：Gin

3. 数据库：MySQL

4. 中间件：跨域中间件、JWT鉴权中间件

5. 架构模式：MVC架构

6. 安全性：数据加密

前端架构

1. 开发语言：HTML、CSS、JavaScript

2. 框架：Vue

3. 工具：
   
   代码编辑器：VS Code
   
   构建工具：Vite
   
   调试工具：Vue.js devtools

4. 组件库：Element Plus

5. 设计模式：Vue Router策略模式

#### 

#### 特技

> 在这么冷的天  
> 想抽根电子烟  
> 可瑞克没有电  
> 可是雪豹已失联  
> 摘不下我​虚伪的假​面  
> 几句胡言被奉为圣箴  
> 尝一口你血汗的香甜  
> 可是钱飘进双眼  
> 模糊了我做人的底线  
> 心甘情愿舍弃了尊严  
> 烧一朵花圈作祭奠  
> 那些说我们的歌土的人，说抄袭的人，说我们现场拉的人，说我饭圈的人，他们这辈子也写不出这样的歌，他们这辈子也听不到你们合唱了。  
> 我是一个特别固执的人，我从来不会在意别人跟我说什么，让我去做。如果你也可以像我一样，那我觉得这件事情——泰裤辣！！