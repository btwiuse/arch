FROM btwiuse/arch:dev

RUN curl -sL https://aka.ms/install-vscode-server/setup.sh | sed -e 's,GLIBC_VERSION=.*$,GLIBC_VERSION=$(ldd --version | head -n1 | grep -oE "[^ ]*" | tail -n1),g' | grep -v chown | bash

RUN pacman -Syu --noconfirm --needed --overwrite='*' yaourt

ENV GOBIN /usr/local/bin
ENV CARGO_INSTALL_ROOT /usr/local

USER btwiuse

# xpanes
RUN yaourt -S --noconfirm tmux-xpanes

WORKDIR /home/btwiuse/

# copilot
RUN sudo pacman -Syu --noconfirm --needed --overwrite='*' neovim
RUN git clone https://github.com/github/copilot.vim.git .config/nvim/pack/github/start/copilot.vim
RUN git clone https://github.com/github/copilot.vim.git .vim/pack/github/start/copilot.vim
RUN mkdir -p .config/github-copilot
# RUN echo '{"github.com":{"user":"zoom399","oauth_token":"ghu_9dzDNqp4A3SojlIQPTaqAqjSbaTjjD1GxYJW"}}' > .config/github-copilot/hosts.json
RUN base64 -d <<< eyJnaXRodWIuY29tIjp7InVzZXIiOiJ6b29tMzk5Iiwib2F1dGhfdG9rZW4iOiJnaHVfVW51dkQ1eTdJWnZDY09nWGtvdWpmUWgzaTBHbFB4MFk0ZUozIn19Cg== > .config/github-copilot/hosts.json

# bash
RUN echo 'set -o vi' >> .bashrc

# code-server
RUN echo 'code-server serve-local --accept-server-license-terms --quality insiders --without-connection-token --port 8000' > code-server
RUN echo 'ufo teleport https://ufo.k0s.io/code http://localhost:8000' >> code-server
