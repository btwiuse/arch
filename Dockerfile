FROM btwiuse/arch:gear AS builder

RUN sudo pacman -S --noconfirm --needed protobuf

RUN git pull && git checkout 06271b3d82154ba728074ff1faeb34a681ad3aee

COPY rust-toolchain ./

RUN env TERM=tmux-256color cargo build --release

FROM btwiuse/arch:k0s

COPY --from=builder /gear/target/release/gear /usr/local/bin/gear
