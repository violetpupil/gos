# [cors](https://www.ruanyifeng.com/blog/2016/04/cors.html)

Cross-origin resource sharing 跨域资源共享

## 简单请求

1. 请求方法 HEAD GET POST
2. 请求头 Accept Accept-Language Content-Language Last-Event-ID Content-Type
3. Content-Type application/x-www-form-urlencoded、multipart/form-data、text/plain

### 流程

请求头用 Origin 表示当前域 scheme://host[:port]
