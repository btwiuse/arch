FROM btwiuse/arch:golang

ENV GOBIN=/usr/local/bin
ENV GOPHERJS_GOROOT=/root/sdk/go1.21.13

RUN go install golang.org/dl/go1.21.13@latest
RUN go1.21.13 download
RUN go install github.com/gopherjs/gopherjs@v1.21.0
