# [swagger](https://swagger.io/)

[文档](https://swagger.io/docs/)

[swagger-ui](https://swagger.io/docs/open-source-tools/swagger-ui/usage/installation/)

[specification](https://swagger.io/docs/specification/about/)

## [swagger-codegen](https://swagger.io/docs/open-source-tools/swagger-codegen/)

[github](https://github.com/swagger-api/swagger-codegen)

```bash
wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.32/swagger-codegen-cli-2.4.32.jar -O swagger-codegen-cli.jar
java -jar swagger-codegen-cli.jar help
# 将.swagger-codegen-ignore文件放在目标目录下，可以只生成模型
java -jar swagger-codegen-cli.jar generate -i https://petstore.swagger.io/v2/swagger.json -l go -o ./swagger-gen
```
