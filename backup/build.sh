#!/bin/bash
os=linux
arch=amd64

# 获取 Docker 镜像标签
get_docker_tag() {
  # 获取当前日期和时间，格式为 YYYY.MM.DD
  date=$(date +%Y.%m.%d)
  # 获取系统信息，格式为 AMD64
  system=$arch
  # 获取当前 Git 仓库的最新提交 ID
  commit=$(git rev-parse HEAD)
  # 组合字符串，格式为 日期-系统-commitId
  tag="$date-$system-$commit"
  echo $tag
}

docker_tag=$(get_docker_tag)

GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o ./temp/linux_amd64/main main.go

docker build -f manifest/docker/Dockerfile -t coding-delivery-docker.pkg.coding.net/coding-ka/tmp-docker/go-example:$docker_tag .

docker push coding-delivery-docker.pkg.coding.net/coding-ka/tmp-docker/go-example:$docker_tag