FROM centos:7.6.1810
MAINTAINER Sushil Lakra

WORKDIR /root/

COPY LinkedListServer .
COPY ./entrypoint.sh .
RUN chmod +x ./entrypoint.sh

ENV PORT 8090
EXPOSE 8090

CMD ["sh","./entrypoint.sh"]

