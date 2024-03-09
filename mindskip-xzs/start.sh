#!/bin/bash

## 设置默认环境变量
export JAVA_HOME=${JAVA_HOME:-/data/jdk1.8}
export PATH=$JAVA_HOME/bin:$PATH
BIN_PATH=`dirname $0`
APP_DIR=${BIN_PATH}/..

# get params
function usage() {
    echo "usage: $0 [PROFILE]"
}

if [ $# -eq 0  ]; then
    usage
    exit 1
fi

PROFILE=$1
shift

APP_JAR=$(ls -d ${APP_DIR}/lib/*.jar)
exec java -Dfile.encoding=utf-8 \
    -Djava.security.egd=file:/dev/./urandom \
    -Dspring.profiles.active=${PROFILE} \
    ${JAVA_OPTS} \
    -jar ${APP_JAR} ${APP_OPTS}
