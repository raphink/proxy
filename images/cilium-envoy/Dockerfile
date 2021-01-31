# syntax=docker/dockerfile:1.1-experimental

# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG COMPILERS_IMAGE=docker.io/cilium/image-compilers:c1ba0665b6f9f012d014a642d9882f7c38bdf365

# Explicitly selecting the amd64 image
FROM --platform=linux/amd64 ${COMPILERS_IMAGE} as builder

ARG BAZEL_BUILD_EXTRA_OPTS
ARG BAZEL_BUILD_OPTS="--jobs=auto --experimental_local_memory_estimate --verbose_failures ${BAZEL_BUILD_EXTRA_OPTS}"

WORKDIR /src

RUN apt-get update
RUN apt-get install -y clang-10 llvm-10-dev clang-format-10 lld-10

RUN --mount=type=bind,readwrite,target=/src \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS} --incompatible_enable_cc_toolchain_resolution --platforms=//bazel/platforms:aarch64_cross --define=cross=aarch64 --linkopt=\"-Wl,--reduce-memory-overheads\"" \
  STRIP=aarch64-linux-gnu-strip \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/arm64

RUN --mount=type=bind,readwrite,target=/src \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS}" \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/amd64

FROM scratch
ARG TARGETPLATFORM
LABEL maintainer="maintainer@cilium.io"
COPY images/cilium-envoy/empty /tmp/empty
COPY --from=builder /out/${TARGETPLATFORM} /
