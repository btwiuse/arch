FROM goreleaser/goreleaser AS goreleaser

FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch go clang tinygo

COPY --from=goreleaser /usr/bin/goreleaser /usr/bin/goreleaser
