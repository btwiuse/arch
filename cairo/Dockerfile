FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' asdf-vm nodejs yarn npm unzip && pacman -Scc --noconfirm
RUN npm i -g bun

RUN asdf plugin add scarb
RUN asdf install scarb latest
RUN asdf set -u scarb latest

RUN asdf plugin add starknet-foundry
RUN asdf install starknet-foundry latest
RUN asdf set -u starknet-foundry latest

RUN curl -sL https://github.com/starknet-io/starknet-devnet/releases/download/v0.8.1/starknet-devnet-x86_64-unknown-linux-gnu.tar.gz | tar -xvzC /usr/local/bin/

RUN curl -L https://github.com/starkware-libs/cairo/releases/download/v2.14.1-dev.0/release-x86_64-unknown-linux-musl.tar.gz | tar -xzC /

RUN echo 'export PATH="/cairo/bin:$PATH"' >> /root/.bashrc
RUN echo 'export PATH="$HOME/.asdf/shims:$PATH"' >> /root/.bashrc
