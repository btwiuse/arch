# Build all btwiuse/arch:<branch> images from ./<branch>/Dockerfile.
#
# The branch list is auto-discovered from the directories that contain a
# Dockerfile, so dropping a new ./<branch>/Dockerfile is all that's needed to
# support a new image -- no edits here.
#
# Build-order dependencies between images live in deps.mk (e.g. `rust: rustup`).
# They are merged with the generic recipe below so that `make -j` always builds
# a parent image before the children that FROM it.

REGISTRY ?= btwiuse/arch
PUSH     ?= 0

# Every directory that has a Dockerfile is a buildable image.
BRANCHES := $(patsubst %/Dockerfile,%,$(wildcard */Dockerfile))

.PHONY: all list $(BRANCHES)

all: $(BRANCHES)

list:
	@printf '%s\n' $(BRANCHES)

# Generic recipe: build (and optionally push) one image, tagged by branch name.
# `main` is also the base image btwiuse/arch == :latest.
$(BRANCHES):
	docker build -t $(REGISTRY):$@ ./$@
	@if [ "$@" = "main" ]; then docker tag $(REGISTRY):main $(REGISTRY):latest; fi
ifeq ($(PUSH),1)
	docker push $(REGISTRY):$@
	@if [ "$@" = "main" ]; then docker push $(REGISTRY):latest; fi
endif

# Dependency edges only (no recipes); merged with the rule above.
include deps.mk
