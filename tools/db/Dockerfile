FROM postgres:latest

WORKDIR /setup

RUN set -x && apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y netcat && \
    rm -rf /var/lib/apt/lists/*

COPY *.sql /util/
