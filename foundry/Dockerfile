FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch
RUN curl -L https://foundry.paradigm.xyz | bash
RUN /root/.foundry/bin/foundryup
