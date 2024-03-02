首先是基本的项目文件初始化，为项目构建做好准备。

`go.mod`: 我们的项目的 Go mod 配置文件，记录依赖库信息。
`Makefile`: 用于控制器构建和部署的 Makefile 文件
`PROJECT`: 用于生成组件的 Kubebuilder 元数据

```
domain: tutorial.kubebuilder.io
layout: go.kubebuilder.io/v3-alpha
projectName: project
repo: tutorial.kubebuilder.io/project
resources:
- group: batch
  kind: CronJob
  version: v1
version: 3-alpha
```

我们还可以在 config/ 目录下获得启动配置。现在，它只包含了在集群上启动控制器所需的 Kustomize YAML 定义，但一旦我们开始编写控制器，它还将包含我们的 CustomResourceDefinitions(CRD) 、RBAC 配置和 WebhookConfigurations 。

config/default 在标准配置中包含 Kustomize base ，它用于启动控制器。

其他每个目录都包含一个不同的配置，重构为自己的基础。

config/manager: 在集群中以 pod 的形式启动控制器

config/rbac: 在自己的账户下运行控制器所需的权限