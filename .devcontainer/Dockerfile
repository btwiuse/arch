FROM btwiuse/arch:yay

RUN runuser -u btwiuse -- bash -lc '\
    yay -Syu --noconfirm curl-http3 && \
    yay -Sc --noconfirm'
