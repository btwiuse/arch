FROM goreleaser/goreleaser AS goreleaser

FROM btwiuse/arch:rust

RUN pacman -Syu --noconfirm --overwrite='*' zig zls

RUN cargo binstall -y cargo-zigbuild --install-path /usr/bin

COPY --from=goreleaser /usr/bin/goreleaser /usr/bin/goreleaser
