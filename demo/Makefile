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

apply:
	kubectl apply -f kube-deployment.yaml
	kubectl apply -f kube-service.yaml

watch_curl:
	export SVC_PORT=$$(kubectl get service zero-demo --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"') ; \
	watch -n0.3 -x curl -s $$(minikube ip):$${SVC_PORT}/ready

watch_get_po:
	watch -n0.5 -x kubectl get po

delete:
	kubectl delete -f kube-deployment.yaml || true
	kubectl delete -f kube-service.yaml