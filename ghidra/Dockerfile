FROM btwiuse/arch:yay

RUN runuser -u btwiuse -- bash -lc '\
    yay -Syu --noconfirm ghidra-git && \
    yay -Sc --noconfirm'
