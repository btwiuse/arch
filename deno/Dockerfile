FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' deno && pacman -Scc --noconfirm
RUN deno upgrade
