FROM injectivelabs/injective-core:v1.16.0-beta.3 AS injectived

FROM btwiuse/arch:rust AS builder

RUN cargo install --root /usr/local cargo-generate --features vendored-openssl
RUN cargo install --root /usr/local cargo-run-script
RUN cargo install --root /usr/local cosmwasm-check

FROM btwiuse/arch:rust

RUN pacman -Sy --noconfirm nodejs yarn npm
RUN npm i -g bun

COPY --from=builder /usr/local/bin/* /usr/local/bin/
COPY --from=injectived /usr/local/bin/* /usr/local/bin/
COPY --from=injectived /lib/libwasmvm.* /lib/

RUN rustup target add wasm32-unknown-unknown

RUN curl -L https://foundry.paradigm.xyz | bash
RUN /root/.foundry/bin/foundryup
