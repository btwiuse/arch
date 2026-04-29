FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker yay aria2 expect && pacman -Scc --noconfirm

RUN runuser -u btwiuse -- bash -lc '\
    yay -S --noconfirm mathematica-light && \
    yay -Sc --noconfirm'