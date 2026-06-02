FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch

RUN curl -sL https://k0s.io/install.sh | bash

RUN cp -v /root/.k0s/bin/k0s /usr/local/bin/
