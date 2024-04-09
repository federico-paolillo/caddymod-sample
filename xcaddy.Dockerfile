FROM golang:1.22.1

COPY bin/xcaddy_0.3.5 /go/xcaddy_0.3.5

WORKDIR /go/my_mod/

ENTRYPOINT [ "/go/xcaddy_0.3.5" ]