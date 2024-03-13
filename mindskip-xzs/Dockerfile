FROM alpine as build

RUN apk update && \
    apk add --update --no-cache curl git wget \
        git bash gcc musl-dev upx

RUN git clone https://github.com/seanly/xzs-mysql.git /code

FROM seanly/toolset:openjdk8u372
USER root

COPY --from=build /code/docker/release/xzs-3.9.0.jar /data/service/lib/
COPY --from=build /code/release/web /data/service/web
COPY start.sh /data/service/bin/

WORKDIR /data/service/

ENTRYPOINT ["bash", "bin/start.sh"]
