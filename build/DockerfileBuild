FROM golang:1.17 as builder

ENV GOPROXY https://goproxy.cn,https://goproxy.io,direct
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /data/imagedesign

COPY . /data/imagedesign

RUN go mod download \
    && go build -o imagedesign .

#================pack=================#
FROM alpine:latest as product

WORKDIR  /app
COPY --from=builder /data/imagedesign/config.ini /app/config.ini
COPY --from=builder /data/imagedesign/imagedesign /app
COPY --from=builder /data/imagedesign/web /app/web

ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
    && chmod +x imagedesign


EXPOSE 80
CMD ["-env","local"]
ENTRYPOINT ["./imagedesign"]
