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
COPY ./relation ./relation
COPY ./relation/cmd/rpc/etc/relation.yaml /app/etc/relation.yaml
RUN go build -ldflags="-s -w" -o /app/relation_rpc ./relation/cmd/rpc


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/etc/relation.yaml /app/etc/relation.yaml
COPY --from=builder /app/relation_rpc /app/relation_rpc

CMD ["./relation_rpc"]
