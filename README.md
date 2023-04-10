# higo-framework
@see https://github.com/dengpju/higo-gin

# 使用
### 下载框架源码
```
点击右侧 Releases -> 点击 Source code
```

### 修改go.mod模块
```
module [模块名称]
```

### 修改env\app.yaml
```
# app config目录
APP_CONFIG: /app/config
```

### 修改env\database.yaml
```
# 数据库
DB:
  Default:
    Driver: "mysql"
    Host: ""
    Port: "3306"
    Database: ""
    Username: ""
    Password: ""
    Charset: "utf8mb4"
    Collation: "utf8mb4_unicode_ci"
    Prefix: ""
    LogMode: true
    MaxIdle: 200
    MaxOpen: 1
    MaxLifetime: 10

Redis:
  Default:
    Host: ""
    Auth: ""
    Port: 6379
    Db: 0
    Pool:
      Min_Connections: 1
      Max_Connections: 10
      Connect_Timeout: 10.0
      Wait_Timeout: 3.0
      Heartbeat: -1
      Max_Idle: 3
      Max_Idle_Time: 60
      Max_Conn_Lifetime: 10
      Wait: true
```

### 修改env\serve.yaml
```
# Http服务器
HTTP_HOST:
  Type: http
  Name: http
  # IP或者域名，如:192.168.42.131:6122或:6122
  Addr: 0.0.0.0:1254
  # 请求超时时间(秒)
  ReadTimeout: 5
  # 响应超时时间(秒)
  WriteTimeout: 10

# Https服务器
HTTPS_HOST:
  Type: http
  Name: https
  # IP或者域名，如:192.168.42.131:6123或:6123
  Addr: 0.0.0.0:1255
  # 请求超时时间(秒)
  ReadTimeout: 5
  # 响应超时时间(秒)
  WriteTimeout: 10

# Websocket服务器
WEBSOCKET_HOST:
  Type: websocket
  Name: websocket
  # IP或者域名，如:192.168.42.131:6125或:6125
  Addr: 0.0.0.0:6125
  # 请求超时时间(秒)
  ReadTimeout: 5
  # 响应超时时间(秒)
  WriteTimeout: 10
```