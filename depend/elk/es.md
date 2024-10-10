# es

配置文件 `/usr/share/elasticsearch/config/elasticsearch.yml`

## 测试curl

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

## 测试golang

```golang
package main

import (
 "log"

 "github.com/elastic/go-elasticsearch/v8"
)

func main() {
 config := elasticsearch.Config{
  Username: "elastic",
  Password: "pass",
 }
 es, _ := elasticsearch.NewClient(config)
 log.Println(es.Info())
}
```

## 索引文档

```golang
func IndexDocument(es *elasticsearch.Client, index string, doc []byte) {
 log := zap.L()

 req := esapi.IndexRequest{
  Index: index,
  Body:  bytes.NewReader(doc),
 }

 res, err := req.Do(context.Background(), es)
 if err != nil {
  log.Error("do req error", zap.Error(err))
  return
 }
 defer res.Body.Close()
 log = log.With(zap.String("status", res.Status()))

 if res.IsError() {
  log.Error("index document error")
 } else {
  var r map[string]any
  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
   log.Error("parse response body error", zap.Error(err))
  } else {
   log.Info("index document result", zap.Any("result", r))
  }
 }
}
```
