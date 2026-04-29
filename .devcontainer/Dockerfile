FROM btwiuse/arch

RUN pacman -Sy --noconfirm expect

COPY . /app

COPY --chmod=0755 /app/activate /app/keygen /usr/bin/