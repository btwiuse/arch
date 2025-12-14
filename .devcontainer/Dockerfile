FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker asdf-vm nodejs yarn npm unzip
RUN npm i -g bun

RUN asdf plugin add scarb
RUN asdf install scarb latest
RUN asdf set -u scarb latest

RUN asdf plugin add starknet-foundry
RUN asdf install starknet-foundry latest
RUN asdf set -u starknet-foundry latest

RUN asdf plugin add starknet-devnet
RUN asdf install starknet-devnet latest
RUN asdf set -u starknet-devnet latest

RUN curl -L https://github.com/starkware-libs/cairo/releases/download/v2.14.1-dev.0/release-x86_64-unknown-linux-musl.tar.gz | tar -xzC /

RUN echo 'export PATH="/cairo/bin:$PATH"' >> /root/.bashrc
RUN echo 'export PATH="$HOME/.asdf/shims:$PATH"' >> /root/.bashrc
