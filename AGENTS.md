# btwiuse/arch image monorepo

All the `btwiuse/arch:<name>` Docker images are built from **one branch
(`main`)**. Each image is just a directory:

```
<name>/Dockerfile      # how to build btwiuse/arch:<name>
<name>/...             # any files that Dockerfile COPY/ADDs from its context
```

`docker build ./<name>` is run with that directory as the build context, so
each image dir is self-contained.

## The 30-second mental model

- **Directory = image.** `rust/Dockerfile` → `btwiuse/arch:rust`.
- **`main` is the base.** `main/Dockerfile` is just `FROM btwiuse/arch`, i.e.
  `btwiuse/arch:latest`. Every other image ultimately `FROM`s it.
- **`deps.mk` records build order**, e.g. `rust: rustup`, `rustup: main`. It is
  **auto-generated** from the `FROM btwiuse/arch:<parent>` lines — you don't
  hand-edit it.
- **The Makefile auto-discovers images** via `wildcard */Dockerfile`. Adding an
  image needs no Makefile edits.

## Build

```bash
make list            # list every buildable image
make rust            # build btwiuse/arch:rust, building rustup -> main first
make all             # build everything, in dependency order
make -j4 all         # parallel; deps.mk keeps parents before children
```

Building any image automatically builds its parent chain first (that's what
`deps.mk` is for). Building `main` also tags/pushes `:latest`.

### Push to Docker Hub

`PUSH=0` by default (build only). `PUSH=1` builds **and pushes every image in
the chain**:

```bash
make rust PUSH=1     # pushes main, rustup, rust
make all PUSH=1
```

Override the registry prefix if needed: `make all REGISTRY=myrepo/arch`.

## Add a new image

```bash
mkdir newtool
cat > newtool/Dockerfile <<'EOF'
FROM btwiuse/arch:rust
RUN ...
EOF
# if the Dockerfile COPY/ADDs local files, drop them in newtool/ too
make deps            # regenerate deps.mk (adds `newtool: rust`)
make newtool         # build it
```

That's it — no Makefile or workflow changes.

## Disable / skip an image

- **Build a subset:** just name them — `make rust cairo PUSH=1`.
- **Exclude from `make all` for one run:** `make all -o xmrig`.
- **Disable permanently:** rename so it stops matching `*/Dockerfile`:
  `git mv xmrig/Dockerfile xmrig/Dockerfile.disabled`, then `make deps`.
  `make` no longer sees it and pushes won't trigger it. Rename back to restore.

## deps.mk (generated — do not hand-edit)

```bash
make deps            # regenerate from the Dockerfiles' FROM lines
make deps-check      # fail if deps.mk is stale (run in CI)
```

`gen-deps.sh` emits an edge `child: parent` for each
`FROM btwiuse/arch:<parent>` line (untagged `FROM btwiuse/arch` => `main`;
builder stages count; non-arch base images are ignored).

## CI (.github/workflows/docker-image.yml)

- **`make deps-check`** runs first, so a stale `deps.mk` fails the build.
- **On push:** a `git diff` finds which `<name>/` dirs changed and builds only
  those (their parents come along via `deps.mk`). Touching shared infra
  (`Makefile`, `deps.mk`, `.github/`) rebuilds everything.
- **Manual run (workflow_dispatch):** pick a `target` input (an image name, or
  `all`). The run is titled `Build <target>` so you can tell runs apart.
- All selected images are built and pushed (`PUSH=1`) after logging in with the
  `DOCKERHUB_USERNAME` / `DOCKERHUB_TOKEN` secrets.

> Note: CI rebuilds a changed image **and its parents**, not its *children*.
> If you change a base like `rustup`, manually trigger the downstream images
> (`rust`, `cosmwasm`, ...) or `make all`.
