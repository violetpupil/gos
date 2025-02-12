# index

Elasticsearch 默认创建 1 个副本

```s
# 创建索引
# 3个分片，1个副本
PUT /books
{
  "settings": {
    "index": {
      "number_of_shards": 3,  
      "number_of_replicas": 1 
    }
  },
  "mappings": {
    "properties": {
      "created_at": {
        "type": "date",
        "format": "epoch_millis"
      }
    }
  }
}
# 删除索引
DELETE /books
```
