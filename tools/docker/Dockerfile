# docker build -t component .
# https://docs.docker.com/engine/reference/builder/
# 使用\分行
# 用.dockerignore指定需要排除的文件，不发送给docker daemon

FROM alpine
# 复制上下文的app文件到容器的/app
COPY app /app

# 运行/app
# exec form 可以接收docker run传递的参数
# ENTRYPOINT ["executable", "param1", "param2"]
#
# shell form 不能接收docker run传递的参数
# ENTRYPOINT command param1 param2
# your ENTRYPOINT will be started as a subcommand of /bin/sh -c, which does not pass signals.
# This means that the executable will not be the container’s PID 1 - and will not receive Unix signals
#
# Only the last ENTRYPOINT instruction in the Dockerfile will have an effect.
ENTRYPOINT /app
