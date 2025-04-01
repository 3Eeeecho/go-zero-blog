# 项目简介

一个基于 go-zero 框架实现的微服务博客系统后端，实现了用户管理、文章管理和标签管理的功能。

项目结构如下:

- app: 所有业务代码包含api、rpc以及mq(消息队列)
  
- pkg: 所有的通用组件,包括error代码和消息,ctxdata,middleware等等
  
- data: 包含项目所依赖的所有中间件产生的数据
  
- deploy: 项目开发时的脚本以及mysql数据结构
  
- doc: 接口的api文档
  
- modd.conf: modd热加载配置文件
  

---

# 技术栈

本项目导入了Gorm库来操作数据库数据,删除了go-zero框架原本的sqlx
- **Gorm**：轻量级 ORM，用于 MySQL 数据操作。
- **MySQL**：存储用户、文章和标签数据。
- **Redis**：缓存 JWT Token 和临时数据。
- **Nginx**：反向代理和负载均衡（可选）。
- **RabbitMQ**：消息队列，用于异步任务（如文章发布通知）。
- **Docker**：容器化服务。
- **Docker Compose**：多容器编排。
- **Modd**：开发时热重载工具。
  

# 快速开始
1. 克隆项目:
   ```bash
   git clone https://github.com/3Eeeecho/go-zero-blog.git
   cd go-zero-blog
2. 登录阿里云容器镜像服务
   ```bash
   docker login --username=<你的用户名> crpi-cxv5pdmggamg66ye.cn-hangzhou.personal.cr.aliyuncs.com
3. 拉取镜像（可选）
   ```bash 
   docker pull crpi-cxv5pdmggamg66ye.cn-hangzhou.personal.cr.aliyuncs.com/eecho_go/go-zero-blog:dev
4. 启动服务
   ```bash
   docker-compose up --build

启动服务后,blog需要一点时间下载go依赖模块,可以输入指令来查看是否运行
   ```bash
   docker logs go-zero-blog
