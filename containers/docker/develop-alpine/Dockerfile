FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/arfblocks/goArf && \
  (cd goArf && make geth) && \
  cp goArf/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /goArf && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]
