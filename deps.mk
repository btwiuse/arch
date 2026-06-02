# btwiuse/arch:<branch> docker image build dependencies.
# Each edge `child: parent` is derived from the FROM btwiuse/arch:<parent>
# lines in ./<child>/Dockerfile (including builder stages).
# `main` is the root base image (FROM btwiuse/arch == latest).
# Consumed by the Makefile to order builds; edges only, no recipes.

main:

apk: main
bazel: main
cairo: main
clojure: main
code-server: main
codespace-btwiuse-arch-qrp54q6r29gx9: main
deno: main
dev: main
dotnet: main
foundry: main
golang: main
grafana: main
jupyter: main
k0s: main
nix: main
node: main
node-16: main
node-18: main
node-20: main
node-22: main
node-24: main
powershell: main
prometheus: main
python: main
qemu: main
ruby: main
rustup: main
swift: main
tea: main
texlive: main
tor: main
uupdump: main
v2ray: main
whois: main
windmill: main
xmrig: main
yaourt: main
yay: main
zig: main

bun: node
dkg: golang
ufo: golang main

rust: rustup
rust-nightly: rustup
solana: rustup main

cosmwasm: rust
gear: rust
ink: rust
near: rust
ninja: rust
rust-goreleaser: rust
stylus: rust

curl3: yay
mathematica-keygen: yay

mathematica: mathematica-keygen
mathematica-light: mathematica-keygen
wolframengine: mathematica-keygen
wljs: wolframengine

elizaos: bun
elizaos-dev: dev

vscode-base: bun node-22
vscode-src: vscode-base
vscode-src-codespace: vscode-base
vscode: vscode-src
vscode-dev: vscode-src
vscode-linux-x64: vscode-src
vscode-linux-x64-ci: vscode-src
vscode-reh-linux-x64: vscode-src
vscode-reh-linux-x64-ci: vscode-src
vscode-reh-web-linux-arm64: vscode-src
vscode-reh-web-linux-arm64-ci: vscode-src
vscode-reh-web-linux-x64: vscode-src
vscode-reh-web-linux-x64-ci: vscode-src
vscode-reh-web-linux-x64-marketplace: vscode-src
vscode-reh-web-linux-x64-open-vsx: vscode-src
vscode-web: vscode-src
vscode-web-ci: vscode-src
vscode-web-min: vscode-src
vscode-web-only: vscode-src
