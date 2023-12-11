# sed

```bash
# 使用正则替换文本 a -> b，并输出
sed 's/a/b/g' tmp.txt
# 写回到文件
sed -i 's/a/b/g' tmp.txt
# 写回到文件，并将原文件备份 tmp.txt.bk
sed -i.bk 's/a/b/g' tmp.txt
```
