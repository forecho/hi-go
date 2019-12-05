# Go 学习笔记

## 安装 Go

### Linux

- 下载源码

```sh
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```

- 处理环境变量

修改 `.bashrc` 或者 `.zshrc` 文件，添加：

```
export PATH=$PATH:/usr/local/go/bin
```

如果需要自定义 `GOPATH` （工作目录）

```
export GOPATH=$HOME/Code/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```



## [包管理器 Dep](https://github.com/golang/dep)

安装

```sh
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```