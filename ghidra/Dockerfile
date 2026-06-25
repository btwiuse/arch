FROM btwiuse/arch:yay

RUN pacman -Syu --noconfirm --overwrite='*' clang && pacman -Scc --noconfirm

RUN runuser -u btwiuse -- bash -lc '\
    yay -Syu --noconfirm ghidra-git && \
    yay -Sc --noconfirm'
