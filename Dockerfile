FROM alpine:latest
ADD build/botboi /usr/bin/botboi
ENTRYPOINT ["botboi"]
