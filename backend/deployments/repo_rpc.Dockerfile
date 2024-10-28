FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ./repo ./repo
COPY ./repo/cmd/rpc/etc/repo.yaml /app/etc/repo.yaml
RUN go build -ldflags="-s -w" -o /app/repo_rpc ./repo/cmd/rpc


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/etc/repo.yaml /app/etc/repo.yaml
COPY --from=builder /app/repo_rpc /app/repo_rpc

CMD ["./repo_rpc"]
