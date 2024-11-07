# 准备

我们采用Docker Compose的方式部署，首先请安装好Docker与Docker Compose。该部分不再赘述，请参考[官方文档](https://docs.docker.com/compose/install/)

新建一个目录，名为GitPulse-deploy，并进入

```bash
mkdir GitPulse-deploy
cd GitPulse-deploy
```

克隆GitPulse的代码

```bash
git clone https://github.com/ShellWen/GitPulse.git
```

# 配置

### 服务配置

复制样例配置文件

```bash
cp ./GitPulse/docs/deploy/example/* ./
```

其中：

* 微服务配置：`files/etc`

* docker-compose配置：`compose-example.yaml`

* nginx配置：`example-nginx.conf`

请根据实际情况对上述文件进行修改，以适配您的需要。

### GitHub Token

Token是用来访问GitHub API的重要参数。保存于`./files/github_env`中

首先请[申请Token](https://github.com/settings/tokens)。

文件内容

```bash
GITHUB_API_TOKEN={Token}
```

请将{Token}替换为自己的Token

# 启动

在GitPulse-deploy目录下运行

```bash
docker-compose up --build -d 
```

建议在首次运行时使用`docker-compose up --build`，以观察可能出现的问题。

日志文件位于`./logs`目录下。

在样例配置中，Prometheus与Jaeger接口并未启用。具体的启用方法请参考 [go-zero 文档](https://go-zero.dev/docs/tutorials/monitor/index) 与 go-zero-looklook 文档 [链路追踪](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/doc/chinese/12-%E9%93%BE%E8%B7%AF%E8%BF%BD%E8%B8%AA.md) 与 [服务监控](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/doc/chinese/13-%E6%9C%8D%E5%8A%A1%E7%9B%91%E6%8E%A7.md)。

我们提供了Prometheus的样例配置文件，位于`GitPulse/deployments/prometheus`目录下，请根据实际情况进行使用。