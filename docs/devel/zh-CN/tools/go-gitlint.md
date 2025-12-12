# go-gitlint 使用指南

## 安装

```bash
$ go install github.com/llorllale/go-gitlint
```

## 配置

### githook: commit-msg配置

```bash
$ cd ${IAM_ROOT}
$ cat > .git/hooks/commit-msg << 'EOF'
#!/bin/sh
$(go env GOPATH)/bin/go-gitlint --msg-file="$1"
EOF
$ chmod +x .git/hooks/commit-msg
```

### .gitlint配置

```bash
$ cd ${IAM_ROOT}
$ cat > .gitlint << 'EOF'
--subject-regex=^((Merge branch.*)|((revert: )?(feat|fix|perf|style|refactor|test|ci|docs|chore)(\(.+\))?: [^A-Z].*[^.]$))
--subject-maxlen=80
--body-regex=^([^\r\n]{0,80}(\r?\n|$))*$
EOF
```

## 运行

```bash
$ cd ${IAM_ROOT}
$ go-gitlint
```
