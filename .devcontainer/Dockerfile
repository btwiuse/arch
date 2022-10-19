FROM btwiuse/arch:ufo AS ufo
FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch

# deno
RUN pacman -Syu --noconfirm --overwrite='*' deno

# python
RUN pacman -Syu --noconfirm --overwrite='*' python

# go
RUN pacman -Syu --noconfirm --overwrite='*' go clang

# rust
RUN pacman -Syu --noconfirm --overwrite='*' rustup clang cmake
RUN rustup default stable

# node
RUN pacman -Syu --noconfirm --overwrite='*' yarn npm nodejs-lts-gallium

# htop dstat
RUN pacman -Syu --noconfirm --overwrite='*' htop dstat

# git
RUN git config --global user.name navigaid
RUN git config --global user.gmail navigaid@gmail.com

# k0s
RUN curl -sL https://k0s.io/install.sh | bash
RUN cp -v /root/.k0s/bin/k0s /usr/local/bin/

# ufo
COPY --from=ufo /usr/local/bin/ufo /usr/local/bin/
