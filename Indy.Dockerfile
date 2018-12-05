FROM ubuntu:xenial

RUN apt-get update && \
    apt-get install -y software-properties-common && \
    apt-get install -y apt-transport-https

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 68DB5E88 && \
    add-apt-repository "deb https://repo.sovrin.org/sdk/deb xenial stable" && \
    apt-get update && \
    apt-get install -y libindy

RUN add-apt-repository "deb https://repo.sovrin.org/sdk/deb xenial master" && \
    apt-get update && \
    apt-get install -y libvcx

RUN add-apt-repository "deb https://repo.sovrin.org/sdk/deb xenial master" && \
    apt-get update && \
    apt-get install -y libnullpay