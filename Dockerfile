FROM alpine:3.6

LABEL maintainer="Ohze JSC <thanhbv@sandinh.net>"

ARG VERSION_TAG=v1.0.1

RUN \
    apk add --no-cache --virtual .build-deps ca-certificates curl && \
    echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf && \
    curl -L https://github.com/ohze/flash-socket-policy-docker/releases/download/$VERSION_TAG/gofsp \
        > /usr/bin/gofsp && \
    chmod +x /usr/bin/gofsp && \
    apk del .build-deps

EXPOSE 843

CMD ["gofsp"]
