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

RUN echo 'export PATH="$HOME/.asdf/shims:$PATH"' >> /root/.bashrc
