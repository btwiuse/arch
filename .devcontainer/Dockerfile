FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch

RUN curl -sL https://aka.ms/install-vscode-server/setup.sh | sed -e 's,GLIBC_VERSION=.*$,GLIBC_VERSION=$(ldd --version | head -n1 | grep -oE "[^ ]*" | tail -n1),g' | grep -v chown | bash

# CMD code-server serve-local --accept-server-license-terms --quality insiders --without-connection-token --port 8000
