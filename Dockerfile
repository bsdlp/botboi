FROM alpine:latest
RUN apk --update upgrade && \
    apk add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
ADD build/botboi /usr/bin/botboi
ENTRYPOINT ["botboi"]
