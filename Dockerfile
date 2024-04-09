FROM golang:1.22.1 AS xcaddy

ENV CADDY_VERSION=v2.7.6

COPY bin/xcaddy_0.3.5 /go/xcaddy_0.3.5

COPY src/ /go/my_mod/

RUN ls /go/my_mod

WORKDIR /go/my_mod

RUN /go/xcaddy_0.3.5 build \ 
  --output /go/my_mod/my_caddy \
  --with github.com/federico-paolillo/caddymod-sample=./

RUN chmod u+x /go/my_mod/my_caddy

FROM caddy:2.7.6

COPY --from=xcaddy /go/my_mod/my_caddy /usr/bin/caddy