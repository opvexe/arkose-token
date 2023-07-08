## 安装 chromedriver

```shell
sudo apt install chromium-chromedriver
```

## 安装Go依赖

```shell
git clone https://github.com/tebeka/selenium.git
cd vendor/
go run init.go --alsologtostderr  --download_browsers --download_latest
```

### 启动服务

```shell
Starting ChromeDriver 114.0.5735.198 (c3029382d11c5f499e4fc317353a43d411a5ce1c-refs/branch-heads/5735@{#1394}) on port 9515
Only local connections are allowed.
Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
ChromeDriver was started successfully.
```