FROM btwiuse/arch:rust AS builder

RUN pacman -Syu --noconfirm --needed binaryen yarn npm deno

RUN npm install --force -g sails-js

RUN cargo install --root /usr/local cargo-generate cargo-gbuild cargo-expand sails-cli

RUN curl -#L https://gear-builds.s3.amazonaws.com/gear-v1.7.1-x86_64-unknown-linux-gnu.tar.xz | tar -xJC /usr/bin

