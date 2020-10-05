FROM debian:stretch-slim

WORKDIR /

COPY /bin/myscheduler /usr/local/bin

CMD ["myscheduler"]
