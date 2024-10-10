# elk

```bash
# 创建es容器
# 用户名 elastic
# 记住密码和token
docker run --name es01 --net elastic -p 9200:9200 -it -m 1GB ^
-v D:/docker/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml ^
docker.elastic.co/elasticsearch/elasticsearch:8.15.2
# 创建kibana容器
docker run --name kib01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.15.2
```
