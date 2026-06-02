FROM btwiuse/ufo AS ufo

FROM btwiuse/arch:bun AS bun

FROM btwiuse/arch:node-22 AS vscode

COPY --from=ufo /usr/bin/ufo /usr/bin/ufo

COPY --from=bun /usr/bin/bun /usr/bin/bunx /usr/bin/

RUN npx -y n 22

RUN pacman -Syu --noconfirm --overwrite='*' gcc libxkbfile unzip wget deno rsync bats quilt github-cli && pacman -Scc --noconfirm

ENV VSCODE_SKIP_NODE_VERSION_CHECK=true