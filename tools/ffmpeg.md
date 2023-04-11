# [ffmpeg](https://ffmpeg.org/)

[文档](https://ffmpeg.org/documentation.html)

```bash
# 将当前文件夹所有png文件转为jpg文件，jpg文件更小
for %a in ("*.png") do ffmpeg -i "%a" "%~na.jpg"
```
