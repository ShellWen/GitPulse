FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

COPY ./package.json /app/package.json
WORKDIR /app

RUN --mount=type=cache,id=corepack,target=/corepack-cache \
		ls -la /corepack-cache

RUN --mount=type=cache,id=corepack,target=/corepack-cache \
    if [ ! -f "/corepack-cache/corepack.tgz" ]; then \
		corepack pnpm --version; \
		mkdir -p /corepack-cache; \
    corepack pack -o /corepack-cache/corepack.tgz; \
    fi

RUN --mount=type=cache,id=corepack,target=/corepack-cache \
    COREPACK_ENABLE_NETWORK=0 corepack install -g /corepack-cache/corepack.tgz

COPY . /app

FROM base AS build

ARG BACKEND_BASEURL="https://api-pulse.shellwen.com"
# Use any value except "" to enable mocks
ARG ENABLE_MOCK=""

ENV VITE_BACKEND_BASEURL=$BACKEND_BASEURL
ENV VITE_ENABLE_MOCK=$ENABLE_MOCK

RUN \
	--mount=type=cache,id=pnpm,target=/pnpm/store \
	COREPACK_ENABLE_STRICT=0 COREPACK_ENABLE_NETWORK=0 corepack pnpm install --frozen-lockfile --prefer-offline
RUN corepack pnpm run build

FROM nginxinc/nginx-unprivileged:stable-alpine-slim
COPY --from=build /app/dist /app
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf
WORKDIR /app
EXPOSE 8080
