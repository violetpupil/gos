# [cors](https://www.ruanyifeng.com/blog/2016/04/cors.html)

Cross-origin resource sharing 跨域资源共享

## 简单请求

1. 请求方法 HEAD GET POST
2. 有限的请求头
3. Content-Type application/x-www-form-urlencoded、multipart/form-data、text/plain

### 流程

请求头用 Origin 表示当前域 scheme://host[:port]

响应头包含 Access-Control-Allow-Origin，表示允许请求

1. 设置为请求头 Origin，允许当前域

2. 设置为 *，允许所有域

## 非简单请求

发送 OPTIONS 预检 (preflight) 请求，请求头包括

1. Origin 表示当前域 scheme://host[:port]

2. Access-Control-Request-Method 跨域请求方法

3. Access-Control-Request-Headers 额外的跨域请求头

预检响应头

1. 包含 Access-Control-Allow-Origin，表示允许请求

2. Access-Control-Allow-Methods 允许的所有请求方法

3. Access-Control-Request-Headers 允许的所有额外请求头

4. Access-Control-Max-Age 预检请求的有效期

预检有效期内的每次跨域请求，都按简单请求流程进行
