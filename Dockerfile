FROM btwiuse/arch:dev

RUN curl -sL https://aka.ms/install-vscode-server/setup.sh | sed -e 's,GLIBC_VERSION=.*$,GLIBC_VERSION=$(ldd --version | head -n1 | grep -oE "[^ ]*" | tail -n1),g' | grep -v chown | bash

RUN pacman -Syu --noconfirm --needed --overwrite='*' yaourt

ENV GOBIN /usr/local/bin
ENV CARGO_INSTALL_ROOT /usr/local

USER btwiuse

RUN yaourt -S --noconfirm tmux-xpanes

WORKDIR /home/btwiuse/

RUN echo 'set -o vi' >> .bashrc

RUN echo 'code-server serve-local --accept-server-license-terms --quality insiders --without-connection-token --port 8000' > code-server
RUN echo 'ufo teleport https://ufo.k0s.io/code http://localhost:8000' >> code-server
