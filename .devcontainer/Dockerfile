FROM btwiuse/ufo AS ufo

FROM btwiuse/arch:bun AS bun

FROM btwiuse/arch:node-22 AS vscode

COPY --from=ufo /usr/bin/ufo /usr/bin/ufo

COPY --from=bun /usr/bin/bun /usr/bin/bunx /usr/bin/

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch gcc libxkbfile unzip wget deno rsync bats quilt github-cli
