FROM btwiuse/arch:golang AS builder

RUN pacman -Syu --noconfirm --overwrite='*'

RUN git clone https://github.com/webteleport/ufo /webteleport/ufo

WORKDIR /webteleport/ufo

RUN go mod tidy

RUN CGO_ENABLED=0 GOBIN=/usr/local/bin go install ./cmd/ufo

FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch

COPY --from=builder /usr/local/bin/ufo /usr/local/bin/
