FROM goreleaser/goreleaser AS goreleaser

FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker go clang tinygo

COPY --from=goreleaser /usr/bin/goreleaser /usr/bin/goreleaser
