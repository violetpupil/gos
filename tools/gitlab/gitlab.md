# [gitlab](https://docs.gitlab.com/ee/)

## [CI/CD](https://docs.gitlab.com/ee/topics/build_your_application.html)

### [配置](https://docs.gitlab.com/ee/ci/yaml/)

项目根目录创建 .gitlab-ci.yml

#### Global keywords

`default` 设置job默认参数

`stages` 定义阶段

#### Job keywords

`image` 运行的docker镜像

`only` 指定运行分支

`script` 执行命令

`stage` job运行阶段

`tags` 选择runner

`variables` 定义script中使用的变量

### [Predefined variables](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)

`CI_COMMIT_REF_NAME` The branch or tag name for which project is built.
