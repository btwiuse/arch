FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch unzip

RUN curl -sL https://bun.sh/install | bash

RUN cp /root/.bun/bin/bun /usr/local/bin/
