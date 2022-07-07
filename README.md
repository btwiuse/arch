# Docker Base Image for Arch Linux

[![Build Status](https://travis-ci.org/archlinux/archlinux-docker.svg?branch=master)](https://travis-ci.org/archlinux/archlinux-docker)
[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/arch.svg)](https://hub.docker.com/r/btwiuse/arch)
[![License](https://img.shields.io/github/license/btwiuse/arch?color=%23000&style=flat-round)](https://github.com/btwiuse/arch/blob/master/LICENSE)

## Goals

* No bloat, only most common tools are added
* Initialize pacman keyrings at build stage, making `pacman -Syu` work out of the box
* Additional package repos:
  - archlinuxcn
  - blackarch
  - btwiuse (for personal use)
  - aur (manually install `yay` or `yaourt` first)

## Usage

### Docker / Dockerfile

```
$ docker run -it btwiuse/arch
$ pacman -Syu yay
```

```
FROM btwiuse/arch

RUN pacman -Syu yay
```

### Kubernetes

```
$ kubectl run -it arch --image=btwiuse/arch
If you don't see a command prompt, try pressing enter.
[root@arch /]# pacman -Syu yay
```

### GitHub Codespaces

On https://github.com/codespaces/new, open repo `btwiuse/arch`, then click "Create codespace"
