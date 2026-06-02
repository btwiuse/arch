FROM btwiuse/arch:yay

RUN runuser -u btwiuse -- bash -lc '\
    yay -S visual-studio-code-bin --noconfirm && \
    yay -Sc --noconfirm'

# CMD code serve-web --accept-server-license-terms --quality insiders --without-connection-token --port 8000
