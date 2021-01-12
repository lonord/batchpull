# batchpull
一个批量更新git仓库的小工具

## 安装

```bash
$ go get github.com/lonord/batchpull
```

## 示例

有一个目录包含几个git仓库

```bash
$ pwd
/path/to/directory
$ ls
repo_foo    repo_bar    repo_other
```

在该目录中执行 `batchpull`

```bash
$ batchpull
[PULL] repo_foo... OK
[PULL] repo_bar... OK
[PULL] repo_other... OK
```

这些仓库已经更新到最新状态

## 许可证

MIT