FROM golang:latest
ENV TZ=Asia/Shanghai LANG="C.UTF-8"
RUN apt-get update && apt-get install -y libnss3 chromium-driver

WORKDIR /workspace
COPY . .
RUN chmod +x /workspace/start.sh
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN go build -o arkosed .
EXPOSE 8080
CMD ["/workspace/start.sh"]
