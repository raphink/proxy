# syntax=docker/dockerfile:1.1-experimental

# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG COMPILERS_IMAGE=quay.io/cilium/cilium-envoy-builder:1dbffdfc472eeaac0dd7b9f056cf03ad8bafeae2@sha256:50ad862fae1e58d3f4ce17f5a78ebc24320701c35125819504449958901fde45

# Explicitly selecting the amd64 image
FROM --platform=linux/amd64 ${COMPILERS_IMAGE} as builder

ARG BAZEL_BUILD_EXTRA_OPTS
ARG BAZEL_BUILD_OPTS="--jobs=auto --experimental_local_memory_estimate --verbose_failures ${BAZEL_BUILD_EXTRA_OPTS}"

WORKDIR /src

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache/arm64 \
  df -h ; du -h -d 1 /tmp/bazel-cache/arm64

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache/amd64 \
  df -h ; du -h -d 1 /tmp/bazel-cache/amd64

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache/arm64 \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS}  --disk_cache=/tmp/bazel-cache/arm64/cas --repository_cache=/tmp/bazel-cache/arm64/repository --incompatible_enable_cc_toolchain_resolution --platforms=//bazel/platforms:aarch64_cross --define=cross=aarch64 --linkopt=\"-Wl,--reduce-memory-overheads\"" \
  STRIP=aarch64-linux-gnu-strip \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/arm64 || (df -h ; du -h /tmp/bazel-cache ; exit 1 )

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache/amd64 \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS} --disk_cache=/tmp/bazel-cache/amd64/cas --repository_cache=/tmp/bazel-cache/amd64/repository" \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/amd64

FROM scratch
ARG TARGETPLATFORM
LABEL maintainer="maintainer@cilium.io"
COPY images/cilium-envoy/empty /tmp/empty
COPY --from=builder /out/${TARGETPLATFORM} /
