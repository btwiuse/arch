FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch yay aria2

RUN su btwiuse -c "yay -S visual-studio-code-bin --noconfirm"

# CMD code serve-web --accept-server-license-terms --quality insiders --without-connection-token --port 8000
