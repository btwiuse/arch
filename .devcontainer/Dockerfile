FROM btwiuse/ufo AS ufo

FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker yay aria2

USER btwiuse

RUN yay -S --noconfirm mathematica

USER root

COPY --from=ufo /usr/bin/ufo /usr/bin/

CMD ufo term