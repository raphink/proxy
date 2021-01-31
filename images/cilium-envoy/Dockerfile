# syntax=docker/dockerfile:1.1-experimental

# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG COMPILERS_IMAGE=quay.io/cilium/cilium-envoy-builder:1dbffdfc472eeaac0dd7b9f056cf03ad8bafeae2@sha256:50ad862fae1e58d3f4ce17f5a78ebc24320701c35125819504449958901fde45

FROM ${COMPILERS_IMAGE} as builder

ARG BAZEL_BUILD_EXTRA_OPTS
ARG BAZEL_BUILD_OPTS="--jobs=auto --experimental_local_memory_estimate --verbose_failures ${BAZEL_BUILD_EXTRA_OPTS}"

RUN make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS} --incompatible_enable_cc_toolchain_resolution --platforms=//bazel/platforms:aarch64_cross --define=cross=aarch64 --linkopt=\"-Wl,--reduce-memory-overheads\"" \
  STRIP=aarch64-linux-gnu-strip \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/arm64

RUN make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS}" \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/amd64

FROM scratch
ARG TARGETPLATFORM
LABEL maintainer="maintainer@cilium.io"
COPY images/cilium-envoy/empty /tmp/empty
COPY --from=builder /out/${TARGETPLATFORM} /
