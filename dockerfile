FROM golang:alpine AS build

LABEL author=starfishsfive@gmail.com
ARG APP_VERSION
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go install golang.org/x/tools/cmd/goimports@latest
WORKDIR /app
COPY . .
RUN go build \
    -mod=vendor \
    -ldflags="-s -w" \
    -o sql2struct main.go
FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
 apk update --no-cache && apk add --no-cache tzdata && \
 ln -snf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone ;

RUN export PATH=$PATH:/go/bin
COPY --from=build /app/sql2struct /bin/sql2struct
COPY --from=build /go/bin/goimports /bin/goimports
COPY --from=build /usr/local/go/bin/gofmt /bin/gofmt
RUN mkdir -p /workspace && \
  chmod +x /bin/sql2struct /bin/goimports /bin/gofmt
WORKDIR /workspace
ENTRYPOINT ["sql2struct"]
