FROM goreleaser/goreleaser AS goreleaser

FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' go clang tinygo && pacman -Scc --noconfirm

COPY --from=goreleaser /usr/bin/goreleaser /usr/bin/goreleaser
