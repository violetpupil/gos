# [air](https://github.com/cosmtrek/air) live reload

go install github.com/cosmtrek/air@latest

## 用法

```bash
air # 使用当前目录 .air.toml 配置文件，文件不存在则使用默认配置
air -c .air.toml
air init # 生成默认配置文件 .air.toml
air -c .air.toml -- -h # 使用 -- 分隔air参数和程序参数
```

[配置](https://github.com/cosmtrek/air/blob/master/air_example.toml)
