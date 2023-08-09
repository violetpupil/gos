# [minikube](https://minikube.sigs.k8s.io/docs/)

[install](https://minikube.sigs.k8s.io/docs/start/)

```bash
# 启动集群
minikube start
# 启动web ui
minikube dashboard
# 打开浏览器访问nginx service
minikube service nginx
# 获取nginx service访问地址
minikube service nginx --url
# List the currently supported addons
minikube addons list
# 启用插件
minikube addons enable metrics-server
# 禁用插件
minikube addons disable metrics-server
# 停止集群
minikube stop
# 删除集群
minikube delete
```
