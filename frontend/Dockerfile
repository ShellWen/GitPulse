FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

FROM base AS build

ARG BACKEND_BASEURL="https://api.gitpulse.shellwen.com"
# Use any value except "" to enable mocks
ARG ENABLE_MOCK=""

ENV VITE_BACKEND_BASEURL=$BACKEND_BASEURL
ENV VITE_ENABLE_MOCK=$ENABLE_MOCK

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM nginxinc/nginx-unprivileged:stable-alpine-slim
COPY --from=build /app/dist /app
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf
WORKDIR /app
EXPOSE 8080
