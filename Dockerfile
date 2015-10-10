FROM gliderlabs/alpine:3.2

MAINTAINER Vinh Nguyen <kurei@axcoto.com>

RUN mkdir /app
ADD client/dist /app
ADD rewl.linux /app/rewl

ENTRYPOINT ["run.sh"]
