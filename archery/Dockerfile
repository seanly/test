FROM hhyo/archery-base:sha-d8159f4

COPY ./rootfs /

RUN apt --fix-broken install  -y libio-socket-ssl-perl && \
    dpkg -i /tmp/percona-toolkit_3.5.7-1.bullseye_amd64.deb && \
    rm -rf /tmp/*.deb
