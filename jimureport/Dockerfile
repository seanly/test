FROM maven:3.9.2-eclipse-temurin-11 as build

COPY  ./settings.xml /root/.m2/
WORKDIR /code 
RUN git clone https://github.com/seanly/JimuReport/ /code

RUN wget http://www.datanucleus.org/downloads/maven2/oracle/ojdbc6/11.2.0.3/ojdbc6-11.2.0.3.jar && \
    mvn install:install-file -DgroupId=com.oracle -DartifactId=ojdbc6 -Dversion=11.2.0.3 \ 
        -Dpackaging=jar -Dfile=ojdbc6-11.2.0.3.jar -Dmaven.repo.local=/root/.m2/repository/

RUN cd jimureport-example; mvn -B -e -U clean package

FROM seanly/toolset:oracle-jdk17

WORKDIR /data/app

RUN mkdir -p /data/app/lib
EXPOSE 8085

COPY --from=build /code/jimureport-example/target/jimureport-example-1.5.jar /data/app/lib

ENTRYPOINT ["bin/start.sh"]