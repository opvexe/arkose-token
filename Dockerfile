FROM golang:latest
ENV TZ=Asia/Shanghai LANG="C.UTF-8"
RUN apk update && apk add --no-cache chromium chromium-chromedriver
COPY chromedriver /usr/local/bin/chromedriver
RUN chmod +x /usr/local/bin/chromedriver
ENV PATH="/usr/local/bin:${PATH}"

WORKDIR /workspace
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN go build -o arkosed .
EXPOSE 8080
CMD ["start.sh"]
