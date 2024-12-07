FROM ubuntu:22.04
LABEL maintainer="Guan.XiangWei"

# 设置镜像的语言 支持中文，否则中文都乱码
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8

ENV DOCKER_LOGS=stdout

# 创建app 文件夹
RUN mkdir -p /app
RUN mkdir -p /app/config
# 将app设置位当前路径
WORKDIR /app

COPY retail .

RUN chmod +x retail

EXPOSE 9090

# CMD 设置启动命令
CMD ["/app/retail"]
