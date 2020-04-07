

# 1. what's this?
基于[yapi](https://github.com/YMFE/yapi)的开放接口开发的上传swagger.json接口文件到yapi接口平台的客户端工具。

# 2. why? 
1. 原生的cli是基于node环境，需要安装node环境
2. 原生的cli不支持传多个swagger.json的文件

# 3. how to use?
1. install 
```
go get -u github.com/dangerous1990/yapi-cli

```
2. 编写yapi-import.json
```json
{
  "type": "swagger",
  "token": "yourtoken",
  "file": "user.swagger.json,login.swagger.json",
  "merge": "good",
  "server": "http://yourip:3000"
}
```
3. 上传upload
- path yapi-import.json所在的文件夹
```
yapi-cli -path yourpath

```



