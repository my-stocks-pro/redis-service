BIN=redis-service

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFLAGS=-ldflags '-w -s'
DOCKER=docker
DOCKERBUILD=$(DOCKER) build
DOCKERRUN=$(DOCKER) run
DOCKERPUSH=$(DOCKER) push
AWS=alex

GODEP=dep
NEWDEP=$(GODEP) ensure

AWSECR=848984447616.dkr.ecr.us-east-1.amazonaws.com/$(BIN)
AWSLOGIN=aws ecr --profile alex get-login --no-include-email --region us-east-1 | sed 's|https://||'

LOGPATH=/Users/alex/go/src/github.com/my-stocks-pro/redis-service/app_log

all: go-build docker-build aws-login docker-push clean

local: go-build docker-build docker-run

go-build:
	@echo "Golang build executable..."
	$(NEWDEP)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) -o $(BIN) $(GOFLAGS) main.go

aws-login:
	`eval $(AWSLOGIN)`

docker-build:
	@echo "Docker build service..."
	$(DOCKERBUILD) --no-cache -t $(BIN) .

docker-push:
	@echo "Push Docker image to AWS ECR..."
	docker tag $(BIN):latest $(AWSECR):latest
	$(DOCKERPUSH) $(AWSECR):latest

docker-run:
	$(DOCKERRUN) \
	--rm \
	-ti \
	--name=$(BIN) \
	-v $(LOGPATH):/app_log \
	-p 9006:9006
	$(BIN)

clean:
	@echo "Clean"
	$(GOCLEAN)
	rm -f $(BINARY)

fclean:
	docker stop $(docker ps -q)
	docker rm $(docker ps -a -q -f status=exited)
	docker rmi $(docker images -f dangling=true -q)

rebuild:
	docker-clean docker-build

docker-clean:
	docker rmi $(BIN)