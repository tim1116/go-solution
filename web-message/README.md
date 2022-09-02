基于net/http+gorm实现一个简易留言板网站

### 安装

```shell
cp config/database.example.ini config/database.ini

# 配置数据库参数

go run main.go
```

### docker

```shell
docker build . -t web-message
docker run -idt -e address=192.168.1.5 -p 8085:8085 web-message
```

### 演示

浏览器访问 http://127.0.0.1:8085/