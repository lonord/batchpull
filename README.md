# batchpull
A tool for batch updating git repositories.

[中文说明](README_zh.md)

## Install

```bash
$ go get github.com/lonord/batchpull
```

## Example

A directory which contains some git repos.

```bash
$ pwd
/path/to/directory
$ ls
repo_foo    repo_bar    repo_other
```

Execute `batchpull` in this directory.

```bash
$ batchpull
[PULL] repo_foo... OK
[PULL] repo_bar... OK
[PULL] repo_other... OK
```

And now these repos is up-to-date.

## License

MIT