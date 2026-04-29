FROM btwiuse/arch:golang

RUN pacman -Sy --noconfirm expect

ADD . /app

RUN cp /app/activate /usr/bin/ && chmod +x /usr/bin/activate

RUN go build -o /usr/bin/keygen /app/keygen.go