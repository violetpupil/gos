# [swagger](https://swagger.io/) api描述文件及工具

[文档](https://swagger.io/docs/)

[swagger-ui](https://swagger.io/docs/open-source-tools/swagger-ui/usage/installation/)

[specification](https://swagger.io/docs/specification/about/)

## [swagger-codegen](https://swagger.io/docs/open-source-tools/swagger-codegen/)

[github](https://github.com/swagger-api/swagger-codegen)

```bash
# v2 版本
wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.32/swagger-codegen-cli-2.4.32.jar -O swagger-codegen-cli.jar
java -jar swagger-codegen-cli.jar help
# v3 版本
wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.43/swagger-codegen-cli-3.0.43.jar -O swagger-codegen-cli.jar
java -jar swagger-codegen-cli.jar --help
# 生成代码
# 将.swagger-codegen-ignore文件放在目标目录下，可以忽略文件
# config.json指定golang包名
java -jar swagger-codegen-cli.jar generate -i https://petstore.swagger.io/v2/swagger.json -l go -o ./swagger-gen -c config.json
```
