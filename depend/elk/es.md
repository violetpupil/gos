# es

配置文件 `/usr/share/elasticsearch/config/elasticsearch.yml`

```bash
export ELASTIC_PASSWORD="your_password"
# 创建索引
# 3个分片，2个副本
curl -u elastic:$ELASTIC_PASSWORD -X PUT "localhost:9200/books?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 3,  
      "number_of_replicas": 2 
    }
  }
}
'
# 删除索引
curl -u elastic:$ELASTIC_PASSWORD -X DELETE "localhost:9200/books?pretty"
```
