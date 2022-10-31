FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y \
    build-essential \
    groff-base \
    libdb-dev \
    libssl-dev \
    libsasl2-dev \
    libicu66 \
    curl \
    jq \
    net-tools \
    vim \
    git

RUN git clone https://github.com/openldap/openldap.git && cd openldap && \
    git checkout refs/tags/OPENLDAP_REL_ENG_2_4_50

RUN cd openldap && \
    ./configure --prefix=/app --with-tls -enable-crypt --with-cyrus-sasl --enable-rwm --enable-null --enable-shell && \
    make && \
    make install

COPY configs/* /configs/

RUN mkdir /app/etc/openldap/slapd.d 
#RUN /app/sbin/slaptest -f /configs/slapd.conf.shell -F /app/etc/openldap/slapd.d; exit 0
