zinx-Golang轻量级TCP服务器框架

### 目录结构
```
.
├── test                    // 测试
│   └── internal
│       └── tcp             // tcp测试
├── ziface                  // interface层
└── znet                    // 实现层
```

### 架构
1. server
   1. 方法
       - 停止服务器：资源回收和状态回执
       - 运行服务器：调用Start()，调用后做阻塞处理，在之间可以做扩展功能
       - 初始化
   2. 属性
       - 监听的IP
       - 监听的端口
2. 简单的链接封装和业务绑定
    1. 方法
       1. 启动链接Start()
       2. 停止链接Stop()
       3. 获取当前链接的conn对象(套接字)
       4. 得到链接🆔
       5. 得到客户端链接的地址和端口
       6. 发送数据的方法Send()
       7. 连接绑定的业务处理方法
   2. 属性
       1. socket tcp套接字
       2. 链接的id
       3. 当前链接的状态(是否已经关闭)
       4. 与当前链接绑定的处理业务方法
       5. 等待连接被动退出的channel
3. 基础router模块
   1. request请求封装
      1. 属性
         1. 连接connection
         2. 请求数据
      2. 方法
         1. 得到当前连接
         2. 得到当前数据
         3. 新建request请求
   2. router模块
      1. 抽象的Router
         1. 处理业务之前的方法
         2. 处理业务的方法
         3. 处理业务之后的方法
      2. 具体的BaseRouter
         1. 处理业务之前的方法
         2. 处理业务的方法
         3. 处理业务之后的方法
   3. 集成router
      1. IServer添加路由功能
      2. server增加router成员属性
      3. Connection绑定一个router成员
      4. Connection调用已经注册的Router
   4. 消息封装
      1. 消息结构体