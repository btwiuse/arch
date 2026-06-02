FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' tor && pacman -Scc --noconfirm

RUN curl -sL https://github.com/anyone-protocol/ator-protocol/releases/download/v0.4.9.11/anon-live-linux-amd64.zip | busybox unzip -d /usr/bin/ -

RUN chmod +x /usr/bin/anon /usr/bin/anon-gencert
