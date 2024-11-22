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
cp ./GitPulse/deployments/files ./

cp ./GitPulse/compose.yaml ./compose.yaml
```

其中：

* 微服务配置：`files/etc`

* docker-compose配置：`compose-example.yaml`



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

样例配置默认启用了所有的微服务，可以根据实际情况修改。如果Elasticsearch启动失败，可能是资源占用过高导致的，请考虑释放资源或关闭Elasticsearch服务。