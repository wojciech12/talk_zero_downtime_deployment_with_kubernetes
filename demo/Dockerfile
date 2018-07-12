FROM ubuntu:18.04

RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*

ADD demo_linux /srv/demo
ADD .missy.yml /.missy.yml

CMD ["/srv/demo"]
