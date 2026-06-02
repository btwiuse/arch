FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' yaourt && yes | pacman -Scc

USER btwiuse

ADD xpra /tmp/xpra
RUN bash -v /tmp/xpra