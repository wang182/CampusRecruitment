GOWF DEMO
===================
go web 开发框架演示


## 运行项目
1. 初始化依赖
```
go mod download -x
```

2. 创建配置
```
cp docs/configs/config.yaml.sample ./config.yaml
```
拷贝后编辑配置文件，修改 secretKey 和  mysql 连接信息。

3. 创建 db
连接到 mysql，创建上一步中配置的数据库。

4. 调试启动
```
make run-serve
```

#### 编译
```
make build
```

编辑后文件将生成到 bin 目录。


## 克隆新项目
```
cp -a CampusRecruitment new-project
cd new-project
rm -rf .git bin
find . -type f -exec sed -i '' -re 's/CampusRecruitment/new-project/g' {} \;
```

