FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker neofetch \
                                            pkgfile \
                                            htop \
                                            python \
                                            go clang tinygo go-tools \
                                            rustup clang cmake \
                                            yarn npm nodejs esbuild deno \
    && pacman -Scc --noconfirm

RUN npm install -g n pnpm bun
RUN pkgfile -u

# git
RUN git config --global user.name navigaid
RUN git config --global user.gmail navigaid@gmail.com

# k0s && ufo
RUN curl -sL https://k0s.io/install.sh | bash
RUN cp -v /root/.k0s/bin/* /usr/local/bin/
