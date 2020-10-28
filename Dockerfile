FROM debian:stretch-slim

WORKDIR /

COPY /bin/custom-scheduler /usr/local/bin

CMD ["custom-scheduler"]
