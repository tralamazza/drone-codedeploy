FROM alpine:3.4

RUN \
  apk update && \
  apk add \
   ca-certificates \
    mailcap && \
  rm -rf /var/cache/apk/*

ADD drone-codedeploy /bin/
ENTRYPOINT ["/bin/drone-codedeploy"]
