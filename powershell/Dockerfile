FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' yay aria2 && pacman -Scc --noconfirm

RUN su btwiuse -c "yay -S powershell-bin --noconfirm"
