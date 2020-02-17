FROM golang:1.13-alpine3.10 as go-dev

# References:
# https://talks.golang.org/2014/organizeio.slide#10

ENV CGO_ENABLED=0

# Setup development environment
RUN apk -q --no-progress --purge --no-cache upgrade \
 && apk --no-progress --purge --no-cache add --upgrade \
    curl \
    libc-dev \
    git \
    gcc \
    libc6-compat \
 && apk -q --no-progress --purge --no-cache upgrade \
 && rm -vrf /var/cache/apk/*

ENTRYPOINT [ "syslogd" ]

CMD [ "-n" ]
