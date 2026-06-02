FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch yaourt && yes | pacman -Scc

USER btwiuse

ADD xpra /tmp/xpra
RUN bash -v /tmp/xpra