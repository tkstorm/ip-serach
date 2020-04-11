## ipsearch
> ip serach 查询是基于 http://ip.taobao.com/ipSearch.html 接口代理获取的数据（有请求频率限制）
> 
> 以前使用ip.cn会出现频次限制，可以基于命令行走淘宝接口查询

### 安装
```
// 安装ipsearch命令工具，以及httpd服务
go get -u -v github.com/lupguo/ipsearch/cmd/...
```

### ipsearch 使用
```
$ ipsearch -h
Usage of ipsearch:
  -debug
    	debug for request response content
  -ip string
    	ip to search, myip is current ip (default "myip")
  -mode string
    	response content mode (json|text) (default "text")
  -proxy string
    	request by proxy, using for debug
  -timeout duration
    	set http request timeout seconds (default 10s)

// 查看出口IP相关信息
$ ipsearch
Ip: 210.21.233.226, Network: 联通, Address: 中国 广东 深圳

// 查看指定IP，并以JSON格式输出
$ ipsearch -ip 118.144.149.206 -mode json
{"addr":"中国 北京 北京","network":"鹏博士","ip":"118.144.149.206"}
```

### ipshttpd

支持ipshttpd部署，相关请求会转发到ipshttpd查询服务器，然后将请求代理转发给淘宝查询IP信息。

```
// http服务
$ ipshttpd -listen 127.0.0.1:8680
2019/08/07 18:36:43 ip search httpd listen on 127.0.0.1:8680
// 请求查询
$ curl localhost:8680
Usage:
	//search current client ip information
	curl localhost:8680/ips

	//search for target ip information
	curl localhost:8680/ips?ip=targetIp

// 通过curl查询
$ curl localhost:8680/ips
{"addr":"中国 广东 深圳","network":"鹏博士","ip":"175.191.11.165"}
$ curl 'localhost:8680/ips?ip=175.190.11.16'
{"addr":"中国 辽宁 大连","network":"鹏博士","ip":"175.190.11.16"}
```

### 原始的Ip查询信息
```
$ curl -XPOST -d 'ip=myip' http://ip.taobao.com/service/getIpInfo2.php
```

### 注意
- 2019-08-07：代码目前版本还比较粗糙，会持续完善！
    - [ ] 代理问题，寻求更好用的代理
    - [x] 程序中一些已知的Bug修复 
- 2019-08-08
    - [x] 修复了客户端请求ipshttpd没有获取到正确IP的问题
    - [x] 修复了ipshttpd的handler处理
    - [x] 新增了版本展示
- 2019-08-13
    - [x] 新增Docker环境支持
