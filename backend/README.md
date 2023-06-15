# backend

#### 项目介绍

#### 安装和使用说明

```
go run main.go
```

#### API文档

暂无

#### 数据库说明

##### 1、 banners

| 序号  | 名称        | 描述  | 类型            | 键   | 为空  | 额外  | 默认值 |
| --- | --------- | --- | ------------- | --- | --- | --- | --- |
| 1   | `id`      |     | varchar(191)  | PRI | NO  |     |     |
| 2   | `img_url` |     | varchar(1024) |     | NO  |     |     |
| 3   | `hre_url` |     | varchar(1024) |     | NO  |     |     |

##### 2、 carts

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `user_id`    |     | varchar(20)         |     | NO  |                |     |
| 6   | `good_id`    |     | varchar(20)         |     | NO  |                |     |

##### 3、 categories

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `name`       |     | varchar(50)         |     | NO  |                |     |
| 6   | `picture`    |     | varchar(1024)       |     | NO  |                |     |

##### 4、 chat_lists

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `me`         |     | varchar(255)        |     | NO  |                |     |
| 6   | `you`        |     | varchar(255)        |     | NO  |                |     |

##### 5、 chats

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `me`         |     | varchar(255)        |     | NO  |                |     |
| 6   | `you`        |     | varchar(255)        |     | NO  |                |     |
| 7   | `type`       |     | varchar(255)        |     | NO  |                |     |
| 8   | `content`    |     | varchar(255)        |     | NO  |                |     |

##### 6、 goods

| 序号  | 名称            | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------- | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`          |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at`  |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at`  |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at`  |     | datetime(3)         | MUL | YES |                |     |
| 5   | `user`        |     | varchar(255)        |     | NO  |                |     |
| 6   | `name`        |     | varchar(50)         |     | NO  |                |     |
| 7   | `picture`     |     | varchar(1024)       |     | NO  |                |     |
| 8   | `price`       |     | float               |     | NO  |                |     |
| 9   | `description` |     | varchar(255)        |     | NO  |                |     |
| 10  | `cate_id`     |     | longtext            |     | YES |                |     |
| 11  | `cata_id`     |     | longtext            |     | YES |                |     |
| 12  | `is_sold`     |     | tinyint(1)          |     | YES |                |     |

##### 7、 images

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `blob`       |     | longblob            |     | YES |                |     |
| 3   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 5   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |

##### 8、 orders

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `good_id`    |     | longtext            |     | YES |                |     |
| 6   | `address_id` |     | longtext            |     | YES |                |     |
| 7   | `user_id`    |     | bigint(20) unsigned |     | YES |                |     |
| 8   | `pay_money`  |     | bigint(20)          |     | YES |                |     |

##### 9、 pictures

| 序号  | 名称         | 描述  | 类型            | 键   | 为空  | 额外  | 默认值 |
| --- | ---------- | --- | ------------- | --- | --- | --- | --- |
| 1   | `good_id`  |     | longtext      |     | YES |     |     |
| 2   | `picture1` |     | varchar(1024) |     | NO  |     |     |
| 3   | `picture2` |     | varchar(1024) |     | YES |     |     |
| 4   | `picture3` |     | varchar(1024) |     | YES |     |     |
| 5   | `picture4` |     | varchar(1024) |     | YES |     |     |
| 6   | `picture5` |     | varchar(1024) |     | YES |     |     |

##### 10、 user_addresses

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `receiver`   |     | varchar(255)        |     | NO  |                |     |
| 6   | `contact`    |     | varchar(255)        |     | NO  |                |     |
| 7   | `address`    |     | varchar(1024)       |     | NO  |                |     |
| 8   | `user_id`    |     | varchar(30)         |     | NO  |                |     |

##### 11、 users

| 序号  | 名称           | 描述  | 类型                  | 键   | 为空  | 额外             | 默认值 |
| --- | ------------ | --- | ------------------- | --- | --- | -------------- | --- |
| 1   | `id`         |     | bigint(20) unsigned | PRI | NO  | auto_increment |     |
| 2   | `created_at` |     | datetime(3)         |     | YES |                |     |
| 3   | `updated_at` |     | datetime(3)         |     | YES |                |     |
| 4   | `deleted_at` |     | datetime(3)         | MUL | YES |                |     |
| 5   | `name`       |     | varchar(30)         |     | NO  |                |     |
| 6   | `telephone`  |     | varchar(30)         | UNI | NO  |                |     |
| 7   | `password`   |     | varchar(255)        |     | YES |                |     |
| 8   | `gender`     |     | varchar(10)         |     | NO  |                |     |
| 9   | `avatar`     |     | varchar(1024)       |     | NO  |                |     |
| 10  | `address_id` |     | varchar(255)        |     | NO  |                |     |

#### 部署说明

##### Render部署

Root Directory

```
backend
```

Build Commad

```
go build -tags netgo -ldflags '-s -w' -o app
```

Start Command

```
./app
```

#### 贡献指南

1. clone 本仓库
2. commit 代码
3. push 代码自动生成 pull request
4. 管理员审查代码并合并
