ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf


# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	fi;


# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@gf gen service

# 在构建之前执行所有的单元测试
test:
	go test -v ./...

# 测试代码覆盖率
coverage:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

# 将测试代码覆盖率上传到 codecov
codecov:
	curl -s https://codecov.io/bash > codecov
	chmod +x codecov
	./codecov -f coverage.out -t e2459475-6317-4023-803e-a94d48e44ec6

# 构建命令
build: test
	go build .

.PHONY: test build

# Build image, deploy image and yaml to current kubectl environment and make port forward to local machine.
.PHONY: start
start:
	@set -e; \
	make image; \
	make deploy; \
	make port;

# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git log -1 --format="%cd.%h" --date=format:"%Y%m%d%H%M%S"))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@gf docker -p -b "-a amd64 -s linux -p temp" -tn $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push:
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_TAG};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}";


