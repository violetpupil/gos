# es

配置文件 `/usr/share/elasticsearch/config/elasticsearch.yml`

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

## 搜索文档

```golang
func SearchDocument(es *elasticsearch.Client, index string) {
 log := zap.L()

 var buf bytes.Buffer
 query := map[string]interface{}{
  "query": map[string]interface{}{
   "match": map[string]interface{}{
    "title": "test",
   },
  },
 }
 if err := json.NewEncoder(&buf).Encode(query); err != nil {
  log.Error("encode query error", zap.Error(err))
  return
 }

 res, err := es.Search(
  es.Search.WithContext(context.Background()),
  es.Search.WithIndex(index),
  es.Search.WithBody(&buf),
  es.Search.WithTrackTotalHits(true),
  es.Search.WithPretty(),
 )
 if err != nil {
  log.Error("search req error", zap.Error(err))
  return
 }
 defer res.Body.Close()
 log = log.With(zap.String("status", res.Status()))

 if res.IsError() {
  var e map[string]interface{}
  if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
   log.Error("decode response error", zap.Error(err))
  } else {
   log.Error("search document error", zap.Any("error", e))
  }
 } else {
  var r SearchResult
  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
   log.Error("decode response error", zap.Error(err))
  } else {
   log.Info("search document result",
    zap.Int("total", int(r.Hits.Total.Value)),
    zap.Duration("cost", time.Duration(r.Took)*time.Millisecond))
   for _, hit := range r.Hits.Hits {
    log.Info("find document", zap.String("id", hit.ID), zap.Any("doc", hit.Source))
   }
  }
 }
}
```
