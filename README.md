

# what's this?
基于https://github.com/YMFE/yapi的开放接口开发的上传swagger.json接口文件到yapi接口平台的客户端工具。

# why? 
1. https://github.com/YMFE/yapi原生的cli需要安装node环境
2. 原生的cli不支持传多个swagger.json的文件

# how to use?
1. install 
```
go get -u https://github.com/dangerous1990/yapi-cli

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



