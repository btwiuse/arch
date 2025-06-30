FROM cosmwasm/optimizer:0.17.0 AS cosmwasm-optimizer

FROM injectivelabs/injective-core:v1.16.0-beta.3 AS injectived

FROM btwiuse/arch:rust

RUN pacman -Sy --noconfirm nodejs yarn npm
RUN npm i -g bun

RUN curl -sL https://github.com/btwiuse/ninja/releases/download/cargo-bin/cargo-bin.tgz | tar --strip=1 -xvzC /usr/local

COPY --from=cosmwasm-optimizer /usr/local/bin/bob /usr/local/bin
COPY --from=cosmwasm-optimizer /usr/local/bin/wasm-opt /usr/local/bin
COPY --from=cosmwasm-optimizer /usr/local/bin/optimize.sh /usr/local/bin/cosmwasm-optimizer
RUN sed -i -e 's,PROJECTDIR="$1",PROJECTDIR="${1:-.}",g' -e 's,exit 1,,g' -e s,/target,./target,g -e s,/bin/ash,/bin/bash,g /usr/local/bin/cosmwasm-optimizer

COPY --from=injectived /usr/local/bin/* /usr/local/bin/
COPY --from=injectived /lib/libwasmvm.* /lib/

RUN rustup target add wasm32-unknown-unknown

RUN curl -L https://foundry.paradigm.xyz | bash
RUN /root/.foundry/bin/foundryup
