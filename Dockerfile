FROM gliderlabs/alpine

MAINTAINER Vinh Nguyen <kurei@axcoto.com>

RUN   mkdir         /app
ADD   client/dist   /app
ADD   rewl.linux    /app/rewl
RUN   chmod +x /app/rewl

#ENTRYPOINT ["/bin/sh"]
CMD ["/app/rewl"]
