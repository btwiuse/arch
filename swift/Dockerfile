FROM btwiuse/arch:yay

RUN runuser -u btwiuse -- bash -lc '\
    yay -S gcc --noconfirm && \
    yay -S swift-bin swift-format swiftly --noconfirm && \
    yay -Sc --noconfirm'
