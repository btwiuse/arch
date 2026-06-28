FROM btwiuse/arch:golang

RUN rm -rf /usr/lib/go

ADD https://github.com/justwasm/go4js/releases/download/go1.27.0-go4js.1/go1.27.0-go4js.1.linux-amd64.min.tar.gz /usr/lib/

ENV GOBIN=/usr/local/bin
