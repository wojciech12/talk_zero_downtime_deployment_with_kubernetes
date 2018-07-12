DOCKER_REGISTRY=docker.io
DOCKER_PROJECT_ID=wojciech11
SERVICE_NAME=zero-kube-demo
DOCKER_IMAGE_NAME=$(DOCKER_PROJECT_ID)/${SERVICE_NAME}
DOCKER_IMAGE_REPO_NAME=$(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME)

build:
	go build .

build_linux:
	GOOS=linux go build -o demo_linux .

docker_build: build_linux
	docker build -t $(DOCKER_IMAGE_NAME) . -f Dockerfile

docker_deploy: docker_build
	docker login --username=$${DOCKERHUB_USER} --password=$${DOCKERHUB_PASS}  && \
	docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_REPO_NAME):1.0.0 && \
	docker push $(DOCKER_IMAGE_REPO_NAME):1.0.0

docker_run_echo:
	@echo "docker run --name $(SERVICE_NAME) -p 8090:8090 -d ${DOCKER_PROJECT_ID}/$(SERVICE_NAME):1.0.0 ";