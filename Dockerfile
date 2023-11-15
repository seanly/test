FROM registry.cn-hangzhou.aliyuncs.com/k8ops/node:14

WORKDIR /code

COPY . /code


RUN --mount=type=cache,target=/code/node_modules \
   bash ./build.sh