# BTW I USE ARCH, in Docker

[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/arch.svg)](https://hub.docker.com/r/btwiuse/arch)
[![License](https://img.shields.io/github/license/btwiuse/arch?color=%23000&style=flat-round)](https://github.com/btwiuse/arch/blob/main/LICENSE)

## Goals

* No bloat, only most common tools are added
* Initialize pacman keyrings at build stage, making `pacman -Syu` work out of the box
* Additional package repos:
  - archlinuxcn
  - blackarch
  - [btwiuse](https://github.com/btwiuse/archpkg) (for personal use)
  - aur (manually install `yay` or `yaourt` first)

## Usage

### Docker / Dockerfile

```
$ docker run -it btwiuse/arch
$ pacman -Syu yay
```

[Example Dockerfile](.devcontainer/Dockerfile)
```
FROM btwiuse/arch

RUN pacman -Syu --noconfirm neofetch
```

### Kubernetes

```
$ kubectl run -it arch --image=btwiuse/arch
If you don't see a command prompt, try pressing enter.
[root@arch /]# pacman -Syu neofetch
```

### GitHub Codespaces

On https://github.com/codespaces/new, open repo `btwiuse/arch`, then click "Create codespace"

![image](https://user-images.githubusercontent.com/54848194/177681631-077891d8-670a-4a34-9f6e-1c69b9fc5916.png)
