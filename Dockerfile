FROM gliderlabs/alpine

MAINTAINER Vinh Nguyen <kurei@axcoto.com>

RUN   mkdir         /app
ADD   client/dist   /app
ADD   rewl.linux    /app/rewl
ADD   run.sh        /app/run.sh
RUN   chmod         +x /app/run.sh

ENTRYPOINT ["/bin/sh"]
CMD ["/app/run.sh"]
