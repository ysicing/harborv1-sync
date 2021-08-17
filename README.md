## Usage

1. 编辑main.go 替换域名 & 用户
2. 构建二进制

```bash
GOOS=linux go build
```

3. 生成auth.json

```json
{
  "new.hub": {
    "username": "", 
    "password": ""
  },
  "old.hub": {
    "username": "",
    "password": ""
  }
}
```

4. 构建镜像

```docker
docker build -t xxxx .
docker push xxxx
```

5. 编辑job.yaml, 替换镜像和hosts

6. 等待任务完成