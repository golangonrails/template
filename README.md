# APP Template
## 后端应用服务器模板

### 版本
go version: go1.15.5

### ENV
```
APP_DIR=`PWD` # 应用工作目录
CONFIG_FILE="config.toml" # 应用配置文件
ENV="DEV"|"STAGE"|"PRO" # 运行环境: 开发/测试/生产
```

### LOCAL RUN
```
# 本机用docker-compose启动服务,如数据库等
# 本机启动环境配置在 docker/ 文件夹里
docker-compose up -d

# 拷贝连接docker-compose里服务器的配置文件到 src/
cp config/config.toml src/      # 也可以编辑配置文件指向自己的服务器
# 或者
ln -s ../config/config.toml src/  # 链接到默认的docker服务器配置文件

# 进入代码目录
cd src

# 第一次运行前,打开go mod,安装依赖
export GO111MODULE=on # 建议放到环境变量配置
go mod init

# 直接运行
go run .

# 编译
go build
```
### Tools
#### bin/g.go 代码生成器,可用于生成数据库迁移文件等
```
  go run bin/g.go       # 查看帮助

  go run bin/g.go migration <name>  # 生成migration文件, 然后编辑migration文件来定义对数据库的修改

  go run bin/g.go seed      <name>  # 生成seed文件, 然后编辑seed文件来定义对数据库初始化数据
```
#### Actions
```
  go run . help     # 查看帮助
```
#### db:migration
```
  # 根据生成的migration文件做数据库迁移 (如果没有数据库将自动创建)
  go run . db:migrate

  # 也可通过编译好的二进制文件对数据库进行迁移
  ./app db:migrate
```

### Docker Image
```
# 需要Docker, 在项目目录执行
docker build . -t app_deploy_image --build-arg ARG_GOPROXY=https://goproxy.cn
```

### Code Format
文件/文件夹/包名: 小写下划线   
类名函数名: 大写开头驼峰   
内部类函数名: 小写开头驼峰   
