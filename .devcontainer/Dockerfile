FROM btwiuse/arch:rustup as anchor

RUN rustup default stable

RUN env CARGO_INSTALL_ROOT=/usr/local cargo install --git https://github.com/project-serum/anchor --tag v0.25.0 anchor-cli --locked

FROM btwiuse/arch

COPY --from=anchor /usr/local/bin/anchor /usr/local/bin/anchor

RUN curl -sSfL https://release.solana.com/v1.11.10/install | bash -v

RUN echo PATH="$HOME/.local/share/solana/install/active_release/bin:$PATH" >> ~/.bashrc

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch
