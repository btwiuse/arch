FROM btwiuse/arch:yay

RUN runuser -u btwiuse -- bash -lc '\
    yay -Syu --noconfirm libsql && \
    yay -Sc --noconfirm'
